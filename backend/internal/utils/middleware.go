package utils

import (
	"context"
	"net/http"

	"github.com/lestrrat-go/jwx/v3/jwk"
	"github.com/lestrrat-go/jwx/v3/jwt"
)

func AuthedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		keyset, err := jwk.Fetch(r.Context(), "http://localhost:5173/api/auth/jwks")

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
