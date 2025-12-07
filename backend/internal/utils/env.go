package utils

import (
	"os"
)

type Env struct {
	DB_CONNECTION_STRING string
	FRONTEND_URL         string
}

func GetEnv() (Env, error) {
	return Env{
		DB_CONNECTION_STRING: os.Getenv("DB_CONNECTION_STRING"),
		FRONTEND_URL:         os.Getenv("FRONTEND_URL"),
	}, nil
}
