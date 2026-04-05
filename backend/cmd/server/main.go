package main

import (
	"group-15/backend/internal/routes"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS configuration: Vital for connecting the local Frontend
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} //the frontend port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}

	r.Use(cors.New(config))

	routes.SetupRoutes(r)

	// The backend will run on port 8080
	log.Println("Starting server on port 8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
