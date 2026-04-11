package main

import (
	"backend/internal/config"
	"backend/internal/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 2. Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not upload file . env, system environment variables will be used")
	}

	// Load the configuration
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Database could not be started: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	// CORS configuration: Vital for connecting the local Frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} //the frontend port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(cors.New(config))

	routes.SetupRoutes(r, db)

	// The backend will run on port 8080
	log.Println("Starting server on port 8080...")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
