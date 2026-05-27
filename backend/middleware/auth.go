package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(401, gin.H{"error": "Not Authorized"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// We read the secret of environment variables
	secretKey := os.Getenv("SUPABASE_JWT_SECRET")
	if secretKey == "" {
		// If not configured, the server is shut down for security
		c.JSON(500, gin.H{"error": "Error of the server configuration"})
		c.Abort()
		return
	}

	// The token is analyzed and validated
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		// Converts the string to []byte
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token or expired"})
		c.Abort()
		return
	}

	// The data inside the token is extracted
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// The user ID comes in the "sub" field
		userID, idExists := claims["sub"].(string)

		if !idExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "The token doesn't contain a user ID"})
			c.Abort()
			return
		}

		// Is saved the UUID in the context of Gin
		c.Set("userID", userID)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error reading the token"})
		c.Abort()
		return
	}

	c.Next()
}
