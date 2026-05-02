package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MasterTypeService interface {
	GetAll() ([]dto.MasterTypeDto, error)
}

type MasterTypeHandler struct {
	masterTypeService MasterTypeService
}

func NewMasterTypeHandler(masterTypeService MasterTypeService) *MasterTypeHandler {
	return &MasterTypeHandler{masterTypeService: masterTypeService}
}

func (h *MasterTypeHandler) GetAll(c *gin.Context) {
	types, err := h.masterTypeService.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error when getting types"})
		return
	}

	// The list is returned with a 200 OK
	c.JSON(http.StatusOK, types)
}
