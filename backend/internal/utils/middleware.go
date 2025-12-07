package utils

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func AuthedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env, err := GetEnv()
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		keyset, err := jwk.Fetch(r.Context(), fmt.Sprintf("%s/api/auth/jwks", env.FRONTEND_URL))
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		token, err := jwt.ParseRequest(r, jwt.WithKeySet(keyset))

		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		userId, exists := token.Subject()

		if !exists {
			http.Error(w, "Invalid JWT", http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyUserID, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
