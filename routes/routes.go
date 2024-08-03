package routes

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/aslam-ep/go-e-commerce/docs"
)

func SetupRoutes(r chi.Router) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(struct {
				Message string `json:"message"`
			}{
				Message: "API Endpoints up and running.",
			})
		})

		// Registering the swagger UI handler
		r.Get("/swagger", httpSwagger.WrapHandler)
	})
}
