package main

import (
	"log"
	"net/http"
	"ranku/internal/routes"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Env loaded")

	r := chi.NewRouter()

	env := utils.GetEnv()

	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{env.FRONTEND_URL},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Mount("/characters", routes.CharactersRouter())
	r.Mount("/votes", routes.VotesRouter())
	r.Mount("/likes", routes.LikesRouter())
	r.Mount("/leaderboard", routes.LeaderboardRouter())

	log.Println("Server listening on 4000")
	http.ListenAndServe(":4000", r)
}
