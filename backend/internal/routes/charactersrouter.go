package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ranku/internal/repositories"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func CharactersRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", searchCharacters)

	return r
}

func searchCharacters(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	query := r.URL.Query().Get("query")

	results := []repositories.SearchCharacterRow{}

	if (len(query)) == 0 {
		render.JSON(w, r, results)
		return
	}

	conn, err := utils.GetDbConnection(ctx)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	q := repositories.New(conn)
	results, err = q.SearchCharacter(ctx, fmt.Sprintf("%%%s%%", query))

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	log.Printf("Search characters returned %d results for query %s", len(results), query)

	if len(results) == 0 {
		render.JSON(w, r, []repositories.SearchCharacterRow{})
		return
	}

	render.JSON(w, r, results)
}
