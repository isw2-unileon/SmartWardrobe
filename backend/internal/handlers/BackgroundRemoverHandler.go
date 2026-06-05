package handlers

import (
	"backend/internal/services"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BackgroundRemoverHandler struct {
	removeBgService *services.RemoveBGService
}

func NewBackgroundRemoverHandler(
	s *services.RemoveBGService,
) *BackgroundRemoverHandler {

	return &BackgroundRemoverHandler{
		removeBgService: s,
	}
}

func (h *BackgroundRemoverHandler) RemoveBackground(
	c *gin.Context,
) {

	file, _, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "image required",
			},
		)
		return
	}

	defer file.Close()

	imageBytes, err := io.ReadAll(file)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": "error reading image",
			},
		)
		return
	}

	result, err := h.removeBgService.RemoveBackground(
		imageBytes,
	)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.Data(
		http.StatusOK,
		"image/png",
		result,
	)
}
