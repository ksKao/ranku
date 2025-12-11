package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"ranku/internal/repositories"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
)

func LeaderboardRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", getLeaderboardHandler)
	r.Get("/live", handleSse)

	return r
}

func getLeaderboardFromRedis() ([]repositories.GetTop100VotedCharactersRow, error) {
	conn, err := utils.GetRedisConnection()

	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	result, err := conn.FTSearchWithArgs(ctx, utils.RedisScoreIndexName, "*", &redis.FTSearchOptions{
		SortBy: []redis.FTSearchSortBy{
			{FieldName: "score", Desc: true},
		},
		Limit: 100,
	}).Result()

	if err != nil {
		return nil, err
	}

	results := []repositories.GetTop100VotedCharactersRow{}

	for _, doc := range result.Docs {
		jsonStr := doc.Fields["$"]

		var character repositories.GetTop100VotedCharactersRow
		err = json.Unmarshal([]byte(jsonStr), &character)

		if err != nil {
			return nil, err
		}

		results = append(results, character)
	}

	return results, nil
}

func getLeaderboard() ([]repositories.GetTop100VotedCharactersRow, error) {
	data, err := getLeaderboardFromRedis()

	if err != nil {
		log.Printf("Get data from redis returned error: %s", err.Error())
	}

	if len(data) != 0 {
		log.Printf("Cache found, returning data...")
		return data, nil
	}

	log.Printf("Failed to get data from Redis.")

	ctx := context.Background()

	dbConn, err := utils.GetDbConnection(ctx)

	if err != nil {
		log.Printf("Error getting DB connection: %s", err.Error())
		return nil, err
	}

	defer dbConn.Close()

	q := repositories.New(dbConn)

	characters, err := q.GetTop100VotedCharacters(ctx)

	if err != nil {
		return nil, err
	}

	return characters, nil
}

func getLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	data, err := getLeaderboard()

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	render.JSON(w, r, data)
}

func handleSse(w http.ResponseWriter, r *http.Request) {
	ch := make(chan int)
	utils.AddChannel(&ch)

	log.Printf("Client connected: %s", r.RemoteAddr)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	defer func() {
		close(ch)
		utils.RemoveChannel(&ch)
		log.Printf("Client closed: %s", r.RemoteAddr)
	}()

	flusher, ok := w.(http.Flusher)
	if !ok {
		log.Println("Could not init http.Flusher")
	}

	for {
		select {
		case message := <-ch:
			log.Println(message)
			data, _ := getLeaderboard()
			if data != nil {
				log.Printf("Sending leaderboard event to %s", r.RemoteAddr)
				jsonStr, err := json.Marshal(data)

				if err != nil {
					log.Println("Failed to marshal data")
					continue
				}

				fmt.Fprintf(w, "data: %s\n\n", string(jsonStr))
				flusher.Flush()
			}
		case <-r.Context().Done():
			log.Printf("Client closed connection: %s", r.RemoteAddr)
			return
		}
	}
}
