package main

import (
	"backend/internal/config"
	"backend/internal/routes"
	"backend/middleware"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Upload the file . env
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not upload file . env, system environment variables will be used")
	}

	// Load the configuration
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Database could not be started: %v", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		defer func() {
			_ = sqlDB.Close()
		}()
	}

	r := gin.Default()

	// CORS configuration: Vital for connecting the local Frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("NEXT_URL")} //the frontend port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	// All calls to the back go through the middleware
	r.Use(cors.New(config))
	r.Use(middleware.AuthMiddleware)

	routes.SetupRoutes(r, db)

	// The backend will run on port 8080
	log.Println("Starting server on port 8080...")
	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
