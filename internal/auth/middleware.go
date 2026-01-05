package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/priyansh7parikh/file-upload-scan/internal/logger"

	"go.uber.org/zap"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	RoleKey   contextKey = "role"
)

func Middleware(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "missing auth header", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid auth header", http.StatusUnauthorized)
				return
			}

			claims, err := ParseToken(parts[1])
			if err != nil {
				logger.Log.Warn("invalid jwt",
					zap.Error(err),
				)
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Role-based authorization
			if requiredRole == "admin" && claims.Role != "admin" {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
