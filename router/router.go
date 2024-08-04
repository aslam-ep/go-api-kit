package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/aslam-ep/go-e-commerce/docs/swagger"
	"github.com/aslam-ep/go-e-commerce/internal/user"
	"github.com/aslam-ep/go-e-commerce/utils"
)

func SetupRoutes(r chi.Router, userHandler *user.UserHandler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteResponse(w, http.StatusAccepted, &struct {
				Message string `json:"message"`
			}{
				Message: "API up and running",
			})
		})

		// Registering the swagger UI handler
		r.Get("/swagger/*", httpSwagger.WrapHandler)

		// Auth Router group
		r.Route("/auth", func(r chi.Router) {
		})

		// User Router group
		r.Route("/users", func(r chi.Router) {
			r.Post("/register", userHandler.CreateUser)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", userHandler.GetUser)
				r.Put("/update", userHandler.UpdateUser)
				r.Put("/password-reset", userHandler.ResetPassword)
				r.Delete("/delete", userHandler.DeleteUser)
			})
		})
	})
}
