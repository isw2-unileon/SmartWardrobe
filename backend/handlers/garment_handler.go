package handlers

import (
	"net/http"

	"backend/services"

	"github.com/gin-gonic/gin"
)

func GetGarments(c *gin.Context) {
	userID := c.Query("user_id")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id requerido"})
		return
	}

	data, err := services.GetUserGarments(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error en BD"})
		return
	}

	c.JSON(http.StatusOK, data)
}
