package handlers

import (
	"backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClipHandler struct {
	clipService *services.ClipService
}

func NewClipHandler(s *services.ClipService) *ClipHandler {
	return &ClipHandler{clipService: s}
}

func (h *ClipHandler) Analyze(c *gin.Context) {

	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "image required",
		})
		return
	}
	defer file.Close()

	response, err := h.clipService.Analyze(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
