package main

import (
	"fmt"
	"log"
	"os"
	"smart-command-center-backend/config"
	"smart-command-center-backend/firebase"
	"smart-command-center-backend/models"
	"smart-command-center-backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	// ENV loading
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Database connection and migration
	config.ConnectDatabase()
	if err := models.Migrate(config.DB); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	if err := models.Seed(config.DB); err != nil {
		log.Fatalf("Seeding failed: %v", err)
	}

	// Firebase initialization
	firebase.InitializeFirebase()

	// Router
	r := routes.SetupRouter()
	r.SetTrustedProxies(nil)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fullURL := fmt.Sprintf("http://localhost:%s", port)
	log.Printf("🚀 Server running at %s\n", fullURL)

	r.Run(":" + port)
}
