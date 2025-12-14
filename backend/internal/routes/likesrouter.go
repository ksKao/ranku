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

func LikesRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(utils.AuthedMiddleware)

	r.Post("/", likeCharacter)
	r.Delete("/", unlikeCharacter)

	return r
}

func likeCharacter(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUser(r)
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	type request struct {
		CharacterId string
	}

	var input request
	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	characterId, err := uuid.Parse(input.CharacterId)

	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	conn, err := utils.GetDbConnection(ctx)
	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	defer conn.Close()

	q := repositories.New(conn)

	liked, err := q.CheckLikeExists(ctx, repositories.CheckLikeExistsParams{
		UserId:      user.ID,
		CharacterId: characterId,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	if liked {
		http.Error(w, "You have already liked this character", http.StatusBadRequest)
		return
	}

	character, _ := q.GetCharacterById(ctx, repositories.GetCharacterByIdParams{
		ID: characterId,
	})

	if len(character) == 0 {
		http.Error(w, "This character does not exist", http.StatusBadRequest)
		return
	}

	err = q.CreateLike(ctx, repositories.CreateLikeParams{
		UserId:      user.ID,
		CharacterId: characterId,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}
}

func unlikeCharacter(w http.ResponseWriter, r *http.Request) {
	user, err := utils.GetUser(r)
	if err != nil {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	type request struct {
		CharacterId string
	}

	var input request
	err = json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	characterId, err := uuid.Parse(input.CharacterId)

	if err != nil {
		http.Error(w, "Invalid character ID", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	conn, err := utils.GetDbConnection(ctx)
	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	defer conn.Close()

	q := repositories.New(conn)

	character, _ := q.GetCharacterById(ctx, repositories.GetCharacterByIdParams{
		ID: characterId,
	})

	if len(character) == 0 {
		http.Error(w, "This character does not exist", http.StatusBadRequest)
		return
	}

	liked, err := q.CheckLikeExists(ctx, repositories.CheckLikeExistsParams{
		UserId:      user.ID,
		CharacterId: characterId,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}

	if !liked {
		http.Error(w, "You have not liked this character yet", http.StatusBadRequest)
		return
	}

	err = q.DeleteLike(ctx, repositories.DeleteLikeParams{
		UserId:      user.ID,
		CharacterId: characterId,
	})

	if err != nil {
		utils.LogError(err)
		utils.WriteGenericInternalServerError(w)
		return
	}
}
