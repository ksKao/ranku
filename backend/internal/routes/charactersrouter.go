package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"ranku/internal/repositories"
	"ranku/internal/utils"
	"slices"
	"strings"

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

	searchResults := []repositories.SearchCharacterRow{}

	if (len(query)) == 0 {
		render.JSON(w, r, searchResults)
		return
	}

	conn, err := utils.GetDbConnection(ctx)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	defer conn.Close()

	q := repositories.New(conn)
	searchResults, err = q.SearchCharacter(ctx, fmt.Sprintf("%s%%", query))

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	log.Printf("Search characters returned %d results for query %s", len(searchResults), query)

	if len(searchResults) == 0 {
		render.JSON(w, r, []repositories.SearchCharacterRow{})
		return
	}

	finalResults := []repositories.SearchCharacterRow{}

	// check for exact match first
	foundIndex := slices.IndexFunc(searchResults, func(char repositories.SearchCharacterRow) bool {
		return strings.EqualFold(char.Name, query)
	})

	if foundIndex != -1 {
		finalResults = append(finalResults, searchResults[foundIndex])
	}

	matchCount := 0
	queryLower := strings.ToLower(query)

	// get 5 name matches
	for _, character := range searchResults {

		if strings.HasPrefix(strings.ToLower(character.Name), queryLower) && !strings.EqualFold(character.Name, queryLower) {
			finalResults = append(finalResults, character)
			matchCount++
		}

		if matchCount == 5 {
			break
		}
	}

	// get 5 anime matches
	matchCount = 0
	for _, character := range searchResults {

		if strings.HasPrefix(strings.ToLower(character.Anime), queryLower) && !strings.HasPrefix(strings.ToLower(character.Name), queryLower) && !strings.EqualFold(character.Name, queryLower) {
			finalResults = append(finalResults, character)
			matchCount++
		}

		if matchCount == 5 {
			break
		}
	}

	render.JSON(w, r, finalResults)
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

	defer conn.Close()

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
