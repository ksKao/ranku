package utils

import (
	"os"
)

type Env struct {
	DB_CONNECTION_STRING    string
	FRONTEND_URL            string
	REDIS_CONNECTION_STRING string
}

func GetEnv() Env {
	return Env{
		DB_CONNECTION_STRING:    os.Getenv("DB_CONNECTION_STRING"),
		FRONTEND_URL:            os.Getenv("FRONTEND_URL"),
		REDIS_CONNECTION_STRING: os.Getenv("REDIS_CONNECTION_STRING"),
	}
}
