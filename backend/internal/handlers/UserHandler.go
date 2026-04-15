package handlers

import (
	"backend/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserServiceInterface interface {
	Login(userName string, password string) (string, error)
}

type UserHandler struct {
	userService UserServiceInterface
}

func NewUserHandler(userService UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

// Login handles the POST /api/login route
func (h *UserHandler) Login(c *gin.Context) {
	var credentials dto.LoginDto

	// Decode JSON
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	// Call the service
	token, err := h.userService.Login(credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Respond successfully
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
