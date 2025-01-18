package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func initMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database("rickmorty").Collection("characters")
}

func getCharacterByName(c *gin.Context) {
    name := c.Param("name")
    var character Character

    filter := bson.M{"name": primitive.Regex{Pattern: name, Options: "i"}}

    err := collection.FindOne(context.TODO(), filter).Decode(&character)
    if err != nil {
        fmt.Println("❌ Error fetching character:", err) 
        c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
        return
    }

    c.JSON(http.StatusOK, character)
}


func main() {
	initMongoDB()

	r := gin.Default()
	r.GET("/character/:name", getCharacterByName)

	r.Run(":8080")
}
