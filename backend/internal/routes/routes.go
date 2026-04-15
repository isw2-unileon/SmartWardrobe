package routes

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

// SetupRoutes set up all endpoints
func SetupRoutes(r *gin.Engine, db *sql.DB) {

	// A. Instantiate the Repository
	userRepo := repository.NewUserRepository(db)

	// B. Instantiate the Service and inject the Repository
	userService := services.NewUserService(userRepo)

	// C. Instantiate the Handler and inject the Service
	userHandler := handlers.NewUserHandler(userService)

	// We group routes under /api for clean URLs
	api := r.Group("/api")
	{
		// Public routes (No JWT required)
		api.POST("/login", userHandler.Login)

	}

}
