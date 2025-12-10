package utils

import (
	"context"
	"fmt"
	"net/http"
	"ranku/internal/repositories"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func TryGetUserIdFromRequest(r *http.Request) (string, bool) {
	env := GetEnv()

	keyset, err := jwk.Fetch(r.Context(), fmt.Sprintf("%s/api/auth/jwks", env.FRONTEND_URL))
	if err != nil {
		return "", false
	}

	token, err := jwt.ParseRequest(r, jwt.WithKeySet(keyset))

	if err != nil {
		return "", false
	}

	return token.Subject()
}

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
