package main

import (
	"log"
	"net/http"
	"ranku/internal/routes"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Mount("/characters", routes.CharactersRouter())

	log.Println("Server listening on 4000")
	http.ListenAndServe(":4000", r)
}
