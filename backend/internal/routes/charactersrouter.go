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
	"github.com/google/uuid"
)

func CharactersRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", searchCharacters)
	r.Get("/{id}", getCharacterById)

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

func getCharacterById(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	id := chi.URLParam(r, "id")

	uuid, err := uuid.Parse(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	conn, err := utils.GetDbConnection(ctx)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	var userIdNullable *string = nil
	userId, exists := utils.TryGetUserIdFromRequest(r)

	if exists {
		userIdNullable = &userId
	}

	q := repositories.New(conn)
	results, err := q.GetCharacterById(ctx, repositories.GetCharacterByIdParams{
		ID:     uuid,
		UserId: userIdNullable,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	if len(results) == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	type response struct {
		repositories.Character
		Animes []string `json:"animes"`
		Likes  int64
		Liked  bool
	}

	firstResult := results[0]

	output := response{
		Character: repositories.Character{
			ID:          firstResult.ID,
			Image:       firstResult.Image,
			Name:        firstResult.Name,
			AnilistId:   firstResult.AnilistId,
			BirthYear:   firstResult.BirthYear,
			BirthMonth:  firstResult.BirthMonth,
			BirthDay:    firstResult.BirthDay,
			BloodType:   firstResult.BloodType,
			Age:         firstResult.Age,
			Description: firstResult.Description,
			Gender:      firstResult.Description,
		},
		Animes: []string{},
		Likes:  firstResult.Likes,
		Liked:  firstResult.Liked,
	}

	for _, character := range results {
		output.Animes = append(output.Animes, character.Anime)
	}

	render.JSON(w, r, output)
}
