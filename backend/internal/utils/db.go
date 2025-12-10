package utils

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

const (
	RedisScoreIndexName = "idx:characters"
)

func GetDbConnection(context.Context) (*pgxpool.Pool, error) {
	ctx := context.Background()
	env := GetEnv()

	conn, err := pgxpool.New(ctx, env.DB_CONNECTION_STRING)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func GetRedisConnection() (*redis.Client, error) {
	env := GetEnv()
	opt, err := redis.ParseURL(env.REDIS_CONNECTION_STRING)

	if err != nil {
		return nil, err
	}

	opt.Protocol = 2
	client := redis.NewClient(opt)

	return client, nil
}
