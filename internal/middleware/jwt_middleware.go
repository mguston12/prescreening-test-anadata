package middleware

import (
	"context"
	"net/http"
	"strings"

	"Screening-Test-Anagata/go-project/pkg/utils"
)

type contextKey string

const UserCtxKey = contextKey("user_id")

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Replace(authHeader, "Bearer ", "", 1)

		claims, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserCtxKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
