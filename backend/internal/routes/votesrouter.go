package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"ranku/internal/repositories"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func VotesRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(utils.AuthedMiddleware)
	r.Post("/", createVote)

	return r
}

func createVote(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUser(r)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	type request struct {
		ForCharacterId     string
		AgainstCharacterId string
	}

	var input request
	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	forUuid, err := uuid.Parse(input.ForCharacterId)

	if err != nil {
		http.Error(w, "Invalid forCharacterId", http.StatusBadRequest)
		return
	}

	againstUuid, err := uuid.Parse(input.AgainstCharacterId)

	if err != nil {
		http.Error(w, "Invalid againstCharacterId", http.StatusBadRequest)
		return
	}

	if input.ForCharacterId == input.AgainstCharacterId {
		http.Error(w, "Cannot vote for and against the same character", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	conn, err := utils.GetDbConnection(ctx)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	q := repositories.New(conn)

	existingVote, _ := q.GetVote(ctx, repositories.GetVoteParams{
		UserId:             user.ID,
		ForCharacterId:     forUuid,
		AgainstCharacterId: againstUuid,
	})

	if existingVote.UserId != "" {
		http.Error(w, "You have already voted for this combination", http.StatusBadRequest)
		return
	}

	err = q.CreateVote(ctx, repositories.CreateVoteParams{
		UserId:             user.ID,
		ForCharacterId:     forUuid,
		AgainstCharacterId: againstUuid,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
