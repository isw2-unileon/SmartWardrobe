package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClothingItemService interface {
	GetAll() ([]dto.ClothingItemDto, error)
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
