package utils

import (
	"os"
)

type Env struct {
	DB_CONNECTION_STRING string
}

func GetEnv() (Env, error) {
	return Env{
		DB_CONNECTION_STRING: os.Getenv("DB_CONNECTION_STRING"),
	}, nil
}
