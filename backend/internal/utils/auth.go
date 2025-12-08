package utils

import (
	"context"
	"net/http"
	"ranku/internal/repositories"
)

func GetUser(r *http.Request) (repositories.User, error) {
	ctx := context.Background()

	conn, err := GetDbConnection(ctx)

	if err != nil {
		return repositories.User{}, err
	}

	q := repositories.New(conn)

	userId := r.Context().Value(KeyUserID).(string)

	return q.GetUserById(ctx, userId)
}
