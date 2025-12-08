package utils

import "log"

func LogError(err error) {
	log.Printf("Error: %s", err.Error())
}
