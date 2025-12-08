package utils

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetDbConnection(context.Context) (*pgxpool.Pool, error) {
	ctx := context.Background()
	env, err := GetEnv()

	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.New(ctx, env.DB_CONNECTION_STRING)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
