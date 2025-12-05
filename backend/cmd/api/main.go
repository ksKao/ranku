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

	r.Mount("/characters", routes.CharactersRouter())

	log.Println("Server listening on 3000")
	http.ListenAndServe(":3000", r)
}
