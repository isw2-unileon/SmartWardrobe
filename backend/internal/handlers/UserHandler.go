package handlers

import (
	"group-15/backend/internal/dto"
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
	var req dto.LoginDto

	// Validate JSON input
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password format"})
		return
	}

	// Call the service to perform the login
	token, err := h.userService.Login(req.UserName, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Return the token to the Frontend
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
