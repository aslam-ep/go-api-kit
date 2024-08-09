package router

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/aslam-ep/go-e-commerce/config"
	// Import for swagger docs for swagger handler
	_ "github.com/aslam-ep/go-e-commerce/docs/swagger"
	"github.com/aslam-ep/go-e-commerce/internal/auth"
	"github.com/aslam-ep/go-e-commerce/internal/user"
	"github.com/aslam-ep/go-e-commerce/router/middleware"
	"github.com/aslam-ep/go-e-commerce/utils"
)

// Router struct to hold router, database and handlers
type Router struct {
	Mux         chi.Router
	apiVersion  string
	authHandler *auth.Handler
	userHandler *user.Handler
}

// NewRouter initialize and setup chi router along with the server
func NewRouter(db *sql.DB) *Router {
	// Initialize router
	r := chi.NewRouter()
	r.Use(chiMiddleware.Logger)
	r.Use(httprate.LimitByIP(config.AppConfig.APIRateLimit, time.Minute))

	// Initialize user domain
	userRepo := user.NewRepository(db)
	userServ := user.NewService(userRepo)
	userHandler := user.NewHandler(userServ)

	// Initialize auth domain
	authRepo := auth.NewRepository(db)
	authServ := auth.NewService(userRepo, authRepo)
	authHandler := auth.NewHandler(authServ)

	return &Router{
		Mux:         r,
		apiVersion:  "/api/v1",
		authHandler: authHandler,
		userHandler: userHandler,
	}
}

// SetupRoutes Initialize end points
func (router Router) SetupRoutes() {
	router.Mux.Route(router.apiVersion, func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			utils.WriteResponse(w, http.StatusAccepted, &utils.MessageRes{
				Success: true,
				Message: "Server up and running.",
			})
		})

		// Registering the swagger UI handler
		r.Get("/swagger/*", httpSwagger.WrapHandler)

		// Auth Router group
		r.Route("/auth", func(r chi.Router) {
			r.Post("/register", router.authHandler.Register)
			r.Post("/login", router.authHandler.Login)
			r.Post("/refresh-token", router.authHandler.RefreshToken)
		})

		// User Router group
		r.With(middleware.AuthMiddleware, middleware.ProfileMiddleware).
			Route("users/{user_id}", func(r chi.Router) {
				r.Get("/", router.userHandler.GetUser)
				r.Put("/update", router.userHandler.UpdateUser)
				r.Put("/reset-password", router.userHandler.ChangePassword)
				r.Delete("/delete", router.userHandler.DeleteUser)
			})
	})
}
