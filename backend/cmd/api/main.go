package main

import (
	"log"
	"net/http"
	"ranku/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Env loaded")

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Mount("/characters", routes.CharactersRouter())

	log.Println("Server listening on 4000")
	http.ListenAndServe(":4000", r)
}
