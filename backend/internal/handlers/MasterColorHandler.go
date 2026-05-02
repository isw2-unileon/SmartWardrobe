package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterColorService interface {
	GetAll() ([]dto.MasterColorDto, error)
}

type MasterColorHandler struct {
	masterColorService MasterColorService
}

func NewMasterColorHandler(masterColorService MasterColorService) *MasterColorHandler {
	return &MasterColorHandler{masterColorService: masterColorService}
}

func (h *MasterColorHandler) GetAll(c *gin.Context) {
	colors, err := h.masterColorService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting colors"})
		return
	}

	// The list is returned with a 200 OK
	c.JSON(http.StatusOK, colors)
}
