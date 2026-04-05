package routes

import (
	"group-15/backend/internal/handlers"
	"group-15/backend/internal/repository"
	"group-15/backend/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura todos los endpoints
func SetupRoutes(r *gin.Engine) {

	// A. Instantiate the Repository (Data layer)
	userRepo := repository.NewUserRepository()

	// B. Instantiate the Service and inject the Repository (Logic layer)
	userService := services.NewUserService(userRepo)

	// C. Instantiate the Handler and inject the Service (Web layer)
	userHandler := handlers.NewUserHandler(userService)
	// We group routes under /api for clean URLs
	api := r.Group("/api")
	{
		// Public routes (No JWT required)
		api.POST("/login", userHandler.Login)

		// In the future you will add more here:
		// api.POST("/register", userHandler.Register)
	}

}
