package main

import (
	"log"
	"github.com/KaratayBerkay/go-gin-backend-project/database"
	"github.com/KaratayBerkay/go-gin-backend-project/routes"
)


func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize router
	r := routes.SetupRouter()

	// Start server with graceful shutdown
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
