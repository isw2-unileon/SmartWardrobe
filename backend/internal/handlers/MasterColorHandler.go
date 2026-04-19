package handlers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterColorHandler struct {
	masterColorService *services.MasterColorService
}

func NewMasterColorHandler(masterColorService *services.MasterColorService) *MasterColorHandler {
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
