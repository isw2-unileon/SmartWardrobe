package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	//Ruta raíz (para probar que el servidor funciona)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Smart Wardrobe API funcionando",
		})
	})

	//Rutas de garments
	r.GET("/garments", middleware.AuthMiddleware, handlers.GetGarments)
}
