package main

import (
	"log"
	"net/http"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/database"
	"github.com/aslam-ep/go-e-commerce/handlers"
	"github.com/aslam-ep/go-e-commerce/repositories"
	"github.com/aslam-ep/go-e-commerce/routes"
	"github.com/aslam-ep/go-e-commerce/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Load configurations
	config.LoadConfig()
	log.Println("Loaded configuration values.")

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected to database.")

	// Initilize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Initilize user domain
	userRepo := repositories.NewUserRepository(db)
	userServ := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userServ)

	// Setup routes
	routes.SetupRoutes(r, userHandler)

	// Start the server
	log.Println("Starting the server on :", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, r); err != nil {
		log.Fatalf("Could not start the server: %v\n", err)
	}
}
