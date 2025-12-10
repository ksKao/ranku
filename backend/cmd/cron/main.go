package main

import (
	"context"
	"fmt"
	"log"
	"ranku/internal/repositories"
	"ranku/internal/services/anilist"
	"ranku/internal/utils"
	"slices"
	"strings"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func updateAnimeCharacterCache() {
	ctx := context.Background()

	conn, err := utils.GetDbConnection(ctx)

	if err != nil {
		log.Printf("Error getting DB connection: %s", err.Error())
		return
	}

	defer conn.Close()

	q := repositories.New(conn)

	// fetch 50 pages of anime
	for i := 1; i <= 50; i++ {
		log.Printf("Fetching page %d", i)

		topAnimes, err := anilist.GetAnilistTopAnimeWithCharacters(i)

		if err != nil {
			log.Printf("Error getting data from Anilist: %s", err.Error())
			continue
		}

		log.Printf("Page %d returns %d result(s)", i, len(topAnimes.Data.Page.Media))

		for _, anime := range topAnimes.Data.Page.Media {
			title := anime.Title.Romaji

			if len(title) == 0 {
				log.Printf("Title is empty for anilist ID %d. Skipping...", anime.Id)
				continue
			}

			log.Printf("Start handling upsert for anilist ID %d.", anime.Id)

			animeFromDb, err := q.GetAnimeByAnidbId(ctx, int32(anime.Id))

			if err != nil {
				log.Printf("Anime wtih Anilist ID %d does not exist in DB. Inserting...", anime.Id)

				animeFromDb, err = q.CreateAnime(ctx, repositories.CreateAnimeParams{
					Name:      title,
					AnilistId: int32(anime.Id),
				})

				if err != nil {
					log.Printf("Unable to create anime for anilist ID %d: %s", anime.Id, err.Error())
					continue
				}
			} else if animeFromDb.Name != title {
				log.Printf("Anime with Anilist ID %d title changed. Updating DB...", anime.Id)

				_, err := q.UpdateAnimeNameById(ctx, repositories.UpdateAnimeNameByIdParams{
					ID:   animeFromDb.ID,
					Name: title,
				})

				if err != nil {
					log.Printf("Unable to update anime for anilist ID %d: %s", anime.Id, err.Error())
					continue
				}
			} else {
				log.Printf("Anime with Anilist ID %d is still the same. Not upserting...", anime.Id)
			}

			for _, character := range anime.Characters.Nodes {
				characterFromDb, err := q.GetCharacterByAnilistId(ctx, int32(character.Id))

				if err != nil {
					log.Printf("Character with Anilist ID %d does not exist in DB. Inserting...", character.Id)

					characterFromDb, err = q.CreateCharacter(ctx, repositories.CreateCharacterParams{
						Name:        character.Name.Full,
						Image:       character.Image.Medium,
						AnilistId:   character.Id,
						BirthYear:   character.DateOfBirth.Year,
						BirthMonth:  character.DateOfBirth.Month,
						BirthDay:    character.DateOfBirth.Day,
						BloodType:   character.BloodType,
						Age:         character.Age,
						Description: character.Description,
						Gender:      character.Gender,
					})

					if err != nil {
						log.Printf("Unable to create character for anilist ID %d: %s", character.Id, err.Error())
						continue
					}
				} else {
					log.Printf("Character with Anilist ID %d exists in DB. Updating...", anime.Id)

					err := q.UpdateCharacterById(ctx, repositories.UpdateCharacterByIdParams{
						Name:        character.Name.Full,
						Image:       character.Image.Medium,
						BirthYear:   character.DateOfBirth.Year,
						BirthMonth:  character.DateOfBirth.Month,
						BirthDay:    character.DateOfBirth.Day,
						BloodType:   character.BloodType,
						Age:         character.Age,
						Description: character.Description,
						Gender:      character.Gender,
						ID:          characterFromDb.ID,
					})

					if err != nil {
						log.Printf("Unable to update character for anilist ID %d: %s", character.Id, err.Error())
						continue
					}
				}

				_, err = q.GetAnimeCharacterRelationByIds(ctx, repositories.GetAnimeCharacterRelationByIdsParams{
					AnimeId:     animeFromDb.ID,
					CharacterId: characterFromDb.ID,
				})

				if err == nil {
					log.Printf("Relation already exists for anime %d and character %d. Skipping...", animeFromDb.AnilistId, characterFromDb.AnilistId)
					continue
				}

				err = q.LinkCharacterToAnime(ctx, repositories.LinkCharacterToAnimeParams{
					AnimeId:     animeFromDb.ID,
					CharacterId: characterFromDb.ID,
				})

				if err != nil {
					log.Printf("Unable to link character %d to anime %d: %s", characterFromDb.AnilistId, animeFromDb.AnilistId, err.Error())
					continue
				}
			}
		}

		log.Printf("Sleeping 5 seconds to avoid rate limit...")
		time.Sleep(5 * time.Second)
	}

	log.Printf("Finished caching anime characters")
}

func updateRedisCache() {
	log.Printf("Start updating redis cache")
	ctx := context.Background()

	dbConn, err := utils.GetDbConnection(ctx)

	if err != nil {
		log.Printf("Error getting DB connection: %s", err.Error())
		return
	}

	defer dbConn.Close()

	redisConn, err := utils.GetRedisConnection()
	if err != nil {
		log.Printf("Error getting Redis connection: %s", err.Error())
		return
	}

	defer redisConn.Close()

	q := repositories.New(dbConn)

	characters, err := q.GetTop100VotedCharacters(ctx)

	if err != nil {
		log.Printf("Failed to get top 100 votes: %s", err.Error())
		return
	}

	_, err = redisConn.FTInfo(ctx, utils.RedisScoreIndexName).Result()

	if err != nil {
		log.Printf("Index does not exist. Creating...")

		_, err := redisConn.FTCreate(ctx, utils.RedisScoreIndexName, &redis.FTCreateOptions{
			OnJSON: true,
			Prefix: []any{"characters:"},
		}, &redis.FieldSchema{
			FieldName: "$.score",
			As:        "score",
			FieldType: redis.SearchFieldTypeNumeric,
			Sortable:  true,
		}).Result()

		if err != nil {
			log.Printf("Failed to create index: %s", err.Error())
			return
		}
	}

	// while not recommended in production, the entire database will not exceed 100 keys, which according to Redis:
	// "Redis running on an entry level laptop can scan a 1 million key database in 40 milliseconds"
	keys, err := redisConn.Keys(ctx, "*").Result()
	if err != nil {
		log.Printf("Failed to fetch keys: %s", err.Error())
		return
	}

	for _, key := range keys {
		// check if key is in top 100 votes, if no, remove it
		hasKey := slices.ContainsFunc(characters, func(char repositories.GetTop100VotedCharactersRow) bool {
			return char.ID.String() == strings.Split(key, ":")[1]
		})

		if !hasKey {
			log.Printf("Character with ID %s fell off, deleting...", key)
			_, err := redisConn.Del(ctx, key).Result()
			if err != nil {
				log.Printf("Failed to delete key: %s", err.Error())
				continue
			}
		}
	}

	for _, character := range characters {
		redisConn.JSONSet(ctx, fmt.Sprintf("characters:%s", character.ID.String()), "$", character)
	}

	log.Printf("Finished updating redis cache")
}

func main() {
	log.Println("Loading env...")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Initializing scheduler...")
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err.Error())
	}

	job, err := scheduler.NewJob(gocron.DailyJob(
		1,
		gocron.NewAtTimes(
			gocron.NewAtTime(0, 0, 0),
		),
	), gocron.NewTask(updateAnimeCharacterCache))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Update anime character cache job %s created", job.ID().String())

	// runs every 5 minutes
	job, err = scheduler.NewJob(gocron.CronJob("*/5 * * * *", false), gocron.NewTask(updateRedisCache))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Update redis cache job %s created", job.ID().String())

	scheduler.Start()

	// prevent the program from exiting
	select {}
}
