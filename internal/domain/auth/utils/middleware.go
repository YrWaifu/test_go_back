package utils

import (
	"context"
	"github.com/YrWaifu/test_go_back/pkg/errors"
	"net/http"
	"strings"
)

type AuthFunc func(ctx context.Context, token string) (string, error)

// ZALUPA
func AuthMiddleware(authFunc AuthFunc) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var token string
			if rawToken := strings.Split(r.Header.Get("Authorization"), " "); len(rawToken) == 2 {
				token = rawToken[1]
			}

			id, err := authFunc(r.Context(), token)
			if err == nil {
				r = r.WithContext(InjectAuth(r.Context(), id))
			}

			next.ServeHTTP(w, r)
		})
	}
}

func AuthRequiredMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, ok := ExtractAuth(r.Context())
			if !ok {
				errors.JSONError(w, errors.ErrorResponse{Errors: "auth required"}, http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
