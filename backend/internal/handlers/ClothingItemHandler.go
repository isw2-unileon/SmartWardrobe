package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClothingItemService interface {
	GetAll() ([]dto.ClothingItemDto, error)
	AddClothingItem(dto.ClothingItemDto, dto.UserDto) (bool, error)
}

type ClothingItemHandler struct {
	service ClothingItemService
}

func NewClothingItemHandler(service ClothingItemService) *ClothingItemHandler {
	return &ClothingItemHandler{service: service}
}

func (h *ClothingItemHandler) GetAll(c *gin.Context) {
	clothes, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting clothes"})
		return
	}

	// The list id returned with a 200 OK
	c.JSON(http.StatusOK, clothes)
}

// Add the clothing item of the cody in the database
func (h *ClothingItemHandler) AddClothingItem(c *gin.Context) {
	// Get the clothing item
	var clothingItem dto.ClothingItemDto
	if err := c.ShouldBindJSON(&clothingItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user
	userRaw, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := userRaw.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing user"})
		return
	}

	user := dto.UserDto{
		ID: userID,
	}

	save, err := h.service.AddClothingItem(clothingItem, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when add clothes"})
		return
	}

	// the boolean is returned with a 200 OK
	c.JSON(http.StatusOK, save)
}
