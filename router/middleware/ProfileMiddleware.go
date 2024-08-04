package middleware

import (
	"net/http"

	"github.com/aslam-ep/go-e-commerce/utils"
	"github.com/go-chi/chi/v5"
)

func ProfileMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Retriving user id from context by auth middleware
		userID, ok := r.Context().Value(UserContextKey).(string)
		if !ok {
			utils.WriterErrorResponse(w, http.StatusUnauthorized, "User not authorized")
			return
		}

		// Retriving the user id from url parameter
		paramID := chi.URLParam(r, "id")
		if userID != paramID {
			utils.WriterErrorResponse(w, http.StatusUnauthorized, "User not authorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}