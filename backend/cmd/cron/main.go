package main

import (
	"context"
	"log"
	"ranku/internal/repositories"
	"ranku/internal/services/anilist"
	"ranku/internal/utils"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func updateDbCache() {
	log.Println("Loading env into code...")
	env, err := utils.GetEnv()

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Env loaded into code")

	ctx := context.Background()
	log.Println("Creating new pgxpool...")
	conn, err := pgxpool.New(ctx, env.DB_CONNECTION_STRING)

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Pgxpool created")

	defer conn.Close()

	q := repositories.New(conn)

	// fetch 50 pages of anime
	for i := 1; i <= 2; i++ {
		log.Printf("Fetching page %d", i)

		topAnimes, err := anilist.GetAnilistTopAnimeWithCharacters(i)

		if err != nil {
			log.Fatal(err.Error())
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
	), gocron.NewTask(updateDbCache))
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Job %s created", job.ID().String())

	scheduler.Start()

	// prevent the program from exiting
	select {}
}
