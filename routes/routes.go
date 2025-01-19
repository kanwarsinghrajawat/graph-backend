package routes

import (
	"github.com/gin-gonic/gin"
	"backend/handlers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/character/:name", handlers.GetCharacterByName)
}
