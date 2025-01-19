package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"backend/config" // ✅ Change from "go-rest-api/config" to "backend/config"
	"backend/models" // ✅ Change from "go-rest-api/models" to "backend/models"
)

// GetCharacterByName searches for characters by name (case-insensitive)
func GetCharacterByName(c *gin.Context) {
	name := c.Param("name")
	var characters []models.Character

	filter := bson.M{"name": bson.M{"$regex": fmt.Sprintf(".*%s.*", name), "$options": "i"}}
	cursor, err := config.DB.Collection("characters").Find(context.TODO(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching characters"})
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var character models.Character
		if err := cursor.Decode(&character); err == nil {
			characters = append(characters, character)
		}
	}

	if len(characters) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}

	c.JSON(http.StatusOK, characters)
}
