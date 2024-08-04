package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/utils"
)

type contextKey string

const UserContextKey = contextKey("user")

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.WriterErrorResponse(w, http.StatusUnauthorized, "Authorizarion header is missing")
			return
		}

		// Bearer token
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			utils.WriterErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Validate token
		claims, err := utils.ValidateToken(tokenStr, config.AppConfig.JWTSecret)
		if err != nil {
			utils.WriterErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Store the user id in context
		ctx := context.WithValue(r.Context(), UserContextKey, claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
