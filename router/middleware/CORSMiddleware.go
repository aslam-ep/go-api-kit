package middleware

import (
	"fmt"
	"net/http"

	"github.com/aslam-ep/go-e-commerce/config"
)

// CORS middleware to handle Cross-Origin Resource Sharing
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Setting CORS headers
		w.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%s:%s", config.AppConfig.Domain, config.AppConfig.ServerPort))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// If the request method is OPTIONS, return status 204 (No Content)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
