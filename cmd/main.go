package main

import (
	"log"
	"net/http"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/database"
	"github.com/aslam-ep/go-e-commerce/internal/user"
	"github.com/aslam-ep/go-e-commerce/router"
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
	userRepo := user.NewUserRepository(db)
	userServ := user.NewUserService(userRepo)
	userHandler := user.NewUserHandler(userServ)

	// Setup routes
	router.SetupRoutes(r, userHandler)

	// Start the server
	log.Println("Starting the server on :", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, r); err != nil {
		log.Fatalf("Could not start the server: %v\n", err)
	}
}
