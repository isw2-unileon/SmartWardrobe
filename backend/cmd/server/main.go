package main

import (
	"backend/config"
	"backend/database"
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8080")
}