package routes

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes set up all endpoints
func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	// A. Instantiate the Repository
	masterColorRepo := repository.NewMasterColorRepository(db)
	masterTypeRepo := repository.NewMasterTypeRepository(db)
	masterStyleRepo := repository.NewMasterStyleRepository(db)
	clothingItemRepo := repository.NewClothingItemRepository(db)

	// B. Instantiate the Service and inject the Repository
	masterColorService := services.NewMasterColorService(masterColorRepo)
	masterTypeService := services.NewMasterTypeService(masterTypeRepo)
	masterStyleService := services.NewMasterStyleService(masterStyleRepo)
	clothingItemService := services.NewClothingItemService(clothingItemRepo)

	// C. Instantiate the Handler and inject the Service
	masterColorHandler := handlers.NewMasterColorHandler(masterColorService)
	masterTypeHandler := handlers.NewMasterTypeHandler(masterTypeService)
	masterStyleHandler := handlers.NewMasterStyleHandler(masterStyleService)
	clothingItemsHandler := handlers.NewClothingItemHandler(clothingItemService)

	// We group routes under /api for clean URLs
	api := r.Group("/api")
	{
		api.GET("/getAllColors", masterColorHandler.GetAll)

		api.GET("/getAllTypes", masterTypeHandler.GetAll)

		api.GET("/getAllStyles", masterStyleHandler.GetAll)

		api.GET("/clothingItems", clothingItemsHandler.GetAll)

		api.POST("/addClothingItem", clothingItemsHandler.AddClothingItem)

	}

}
