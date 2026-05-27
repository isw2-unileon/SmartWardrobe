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
		c.JSON(401, gin.H{"error": "No autorizado"})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Leemos el secreto de las variables de entorno de tu ordenador o servidor
	secretKey := os.Getenv("SUPABASE_JWT_SECRET")
	if secretKey == "" {
		// Si se nos olvidó configurarlo, paramos el servidor por seguridad
		c.JSON(500, gin.H{"error": "Error de configuración del servidor"})
		c.Abort()
		return
	}

	// 1. Analizamos y validamos el token usando el secreto de Supabase
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado")
		}
		// Convertimos el string a []byte que es lo que pide la librería
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido o expirado"})
		c.Abort()
		return
	}

	// 2. Extraemos los datos (claims) que Supabase metió dentro del token
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// En Supabase, el ID del usuario SIEMPRE viene en el campo "sub"
		userID, idExists := claims["sub"].(string)

		if !idExists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "El token no contiene un ID de usuario"})
			c.Abort()
			return
		}

		// 3. Guardamos el UUID que nos dio Supabase en el contexto de Gin
		c.Set("userID", userID)
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Error al leer el token"})
		c.Abort()
		return
	}

	c.Next()
}
