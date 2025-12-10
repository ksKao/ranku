package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/redis/go-redis/v9"
)

type redisCharacter struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Image   string `json:"image"`
	For     int    `json:"for"`
	Against int    `json:"against"`
	Score   int    `json:"score"`
}

func LeaderboardRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", getLeaderboard)

	return r
}

func getDataFromRedis() ([]redisCharacter, error) {
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

	results := []redisCharacter{}

	for _, doc := range result.Docs {
		jsonStr := doc.Fields["$"]

		var character redisCharacter
		err = json.Unmarshal([]byte(jsonStr), &character)

		if err != nil {
			return nil, err
		}

		results = append(results, character)
	}

	return results, nil
}

func getLeaderboard(w http.ResponseWriter, r *http.Request) {
	data, err := getDataFromRedis()

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	render.JSON(w, r, data)
}
