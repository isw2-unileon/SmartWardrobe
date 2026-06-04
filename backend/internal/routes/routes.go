package routes

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes set up all endpoints
// func SetupRoutes(r *gin.Engine, db *gorm.DB, clipSvc *services.ClipService) {
func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	// A. Instantiate the Repository
	masterColorRepo := repository.NewMasterColorRepository(db)
	masterTypeRepo := repository.NewMasterTypeRepository(db)
	masterStyleRepo := repository.NewMasterStyleRepository(db)
	masterCategoryRepository := repository.NewMasterCategoriesRepository(db)
	clothingItemRepo := repository.NewClothingItemRepository(db)

	// B. Instantiate the Service and inject the Repository
	masterColorService := services.NewMasterColorService(masterColorRepo)
	masterTypeService := services.NewMasterTypeService(masterTypeRepo)
	masterStyleService := services.NewMasterStyleService(masterStyleRepo)
	masterCategoryService := services.NewMasterCategoriesService(masterCategoryRepository)
	clothingItemService := services.NewClothingItemService(clothingItemRepo)
	locationService := services.NewLocationService()
	weatherService := services.NewWeatherService()
	outfitService := services.NewOutfitService(locationService, weatherService, clothingItemService, masterCategoryService, masterTypeService)

	// C. Instantiate the Handler and inject the Service
	masterColorHandler := handlers.NewMasterColorHandler(masterColorService)
	masterTypeHandler := handlers.NewMasterTypeHandler(masterTypeService)
	masterStyleHandler := handlers.NewMasterStyleHandler(masterStyleService)
	//clipHandler := handlers.NewClipHandler(clipSvc)
	clothingItemsHandler := handlers.NewClothingItemHandler(clothingItemService)
	outfitHandler := handlers.NewOutfitHandler(outfitService)

	// We group routes under /api for clean URLs
	api := r.Group("/api")
	{
		api.GET("/getAllColors", masterColorHandler.GetAll)

		api.GET("/getAllTypes", masterTypeHandler.GetAll)

		api.GET("/getAllStyles", masterStyleHandler.GetAll)

		// CLIP
		//api.POST("/clothing/analyze", clipHandler.Analyze)

		api.GET("/clothingItems", clothingItemsHandler.GetAll)

		api.GET("/clothingItem/filters", clothingItemsHandler.GetClothingItem)

		api.POST("/clothingItem", clothingItemsHandler.AddClothingItem)

		api.GET("/clothingItem/:id", clothingItemsHandler.GetByID)

		api.PUT("/clothingItem/:id", clothingItemsHandler.UpdateClothingItem)

		api.DELETE("/clothingItem/:id", clothingItemsHandler.DeleteClothingItem)

		api.POST("/generateOutfit/days", outfitHandler.GenerateOutfit)

	}

}
