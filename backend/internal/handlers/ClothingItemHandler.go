package handlers

import (
	"backend/internal/dto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClothingItemService interface {
	GetAll() ([]dto.ClothingItemDto, error)
	AddClothingItem(dto.ClothingItemDto, dto.UserDto) (bool, error)
	UpdateClothingItem(int64, dto.ClothingItemDto) (dto.ClothingItemDto, error)
	DeleteClothingItem(int64) error
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

// Update the clothing item that has the id passed as a parameter and updates with the data of the request body
func (h *ClothingItemHandler) UpdateClothingItem(c *gin.Context) {
	//The ID of the clothing item to update is in the URL
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Body with the update data
	var body dto.ClothingItemDto
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.service.UpdateClothingItem(id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating clothing item"})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// Delete the clothing item that has the id passed as a parameter
func (h *ClothingItemHandler) DeleteClothingItem(c *gin.Context) {
	// The ID of the clothing item to delete is in the URL
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.service.DeleteClothingItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting clothing item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
