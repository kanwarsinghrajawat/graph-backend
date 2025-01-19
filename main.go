package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"backend/config"
	"backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found, using system environment variables.")
	}

	config.InitMongoDB()

	router := gin.Default()
	router.Use(cors.Default())

	routes.RegisterRoutes(router)

	port := ":8080"
	fmt.Println("🚀 Server running on http://localhost" + port)

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
