package routes

import (
	"encoding/json"
	"net/http"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
)

type character struct {
	Name string `json:"name"`
}

func CharactersRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(utils.AuthedMiddleware)
	r.Get("/", getAllCharacters)

	return r
}

func getAllCharacters(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	characters := []character{
		{Name: ctx.Value(utils.KeyUserID).(string)},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(characters); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
