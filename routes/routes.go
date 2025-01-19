package routes

import (
	"github.com/gin-gonic/gin"
	"backend/handlers" // ✅ Change "graph-backend" to "backend"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/character/:name", handlers.GetCharacterByName)
}
