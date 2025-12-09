package utils

import (
	"context"
	"net/http"
)

func AuthedMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, exists := TryGetUserIdFromRequest(r)

		if !exists {
			http.Error(w, "", http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, KeyUserID, userId)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
