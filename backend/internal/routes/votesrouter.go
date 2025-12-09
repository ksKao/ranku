package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"ranku/internal/repositories"
	"ranku/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func VotesRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(utils.AuthedMiddleware)

	r.Post("/", createVote)
	r.Get("/matchup", getVoteMatchup)

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

	existingVote, _ := q.GetUserVoteWithCharacterIds(ctx, repositories.GetUserVoteWithCharacterIdsParams{
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

func getVoteMatchup(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUser(r)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
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

	votes, err := q.GetUserVotes(ctx, user.ID)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	allCharacters, err := q.GetAllCharactersByRandomOrder(ctx)

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	// check if user's has exhausted all possible vote combinations
	numberOfCombinations := (len(allCharacters) * (len(allCharacters) - 1)) / 2

	type response struct {
		Char1 *repositories.Character `json:"char1"`
		Char2 *repositories.Character `json:"char2"`
	}

	output := response{
		Char1: nil,
		Char2: nil,
	}

	if numberOfCombinations <= len(votes) {
		render.JSON(w, r, output)
		return
	}

	votedPairs := make(map[string]bool)

	for _, vote := range votes {
		// Store both directions of the matchup (because a vote for A vs B is the same as B vs A)
		votedPairs[fmt.Sprintf("%s-%s", vote.ForCharacterId.String(), vote.AgainstCharacterId.String())] = true
		votedPairs[fmt.Sprintf("%s-%s", vote.AgainstCharacterId.String(), vote.ForCharacterId.String())] = true
	}

	// Find a valid matchup that hasn't been voted on yet
	var char1, char2 *repositories.Character
	var validMatchupFound bool

	for i := 0; i < len(allCharacters)-1; i++ {
		char1 = &allCharacters[i]
		for j := i + 1; j < len(allCharacters); j++ {
			char2 = &allCharacters[j]

			// Ensure the characters are not the same
			if char1.ID == char2.ID {
				continue
			}

			// Check if this combination has been voted on before
			pairID := fmt.Sprintf("%s-%s", char1.ID, char2.ID)
			if _, exists := votedPairs[pairID]; exists {
				continue // This combination was voted on already
			}

			// We found a valid matchup
			validMatchupFound = true
			break
		}
		if validMatchupFound {
			break
		}
	}

	if !validMatchupFound {
		render.JSON(w, r, output)
		return
	}

	output.Char1 = char1
	output.Char2 = char2

	render.JSON(w, r, output)
}
