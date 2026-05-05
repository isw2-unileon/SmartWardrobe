package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterStyleService interface {
	GetAll() ([]dto.MasterStyleDto, error)
}

type MasterStyleHandler struct {
	masterStyleService MasterStyleService
}

func NewMasterStyleHandler(masterStyleService MasterStyleService) *MasterStyleHandler {
	return &MasterStyleHandler{masterStyleService: masterStyleService}
}

func (h *MasterStyleHandler) GetAll(c *gin.Context) {
	colors, err := h.masterStyleService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting styles"})
		return
	}

	// The list is returned with a 200 OK
	c.JSON(http.StatusOK, colors)
}
