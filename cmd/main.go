package main

import (
	"log"
	"net/http"

	"github.com/aslam-ep/go-e-commerce/config"
	"github.com/aslam-ep/go-e-commerce/database"
	"github.com/aslam-ep/go-e-commerce/router"
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

	router := router.NewRouter(db)
	router.SetupRoutes()

	// Start the server
	log.Println("Starting the server on :", config.AppConfig.ServerPort)
	if err := http.ListenAndServe(":"+config.AppConfig.ServerPort, router.Mux); err != nil {
		log.Fatalf("Could not start the server: :%v\n", err)
	}
}
