package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OutfitService interface {
	GenerateOutfit(dto.OutfitRequestDto, dto.UserDto) (*dto.OutfitResponseDto, error)
}

type OutfitHandler struct {
	service OutfitService
}

func NewOutfitHandler(service OutfitService) *OutfitHandler {
	return &OutfitHandler{service: service}
}

// Get the clothing item in function of filters
func (h *OutfitHandler) GenerateOutfit(c *gin.Context) {
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

	// Body with the days and the locate to generate outfit
	var request dto.OutfitRequestDto
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	list, err := h.service.GenerateOutfit(request, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when filter clothes."})
		return
	}

	// the list of outfits with the filters is returned with a 200 OK
	c.JSON(http.StatusOK, list)
}
