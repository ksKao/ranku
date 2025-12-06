package main

import (
	"fmt"
	"log"
	"ranku/internal/services/anilist"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	// fetch 50 pages of anime
	for i := 1; i <= 2; i++ {
		anime, err := anilist.GetAnilistTopAnimeWithCharacters(i)

		if err != nil {
			log.Fatal(err.Error())
		}

		for _, anime := range anime.Data.Page.Media {
			fmt.Println(anime.Characters.Nodes[0].Description)
		}
	}
}
