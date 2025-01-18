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

// ✅ Character struct matching MongoDB schema
type Character struct {
    ID       int      `json:"id" bson:"id"`
    Name     string   `json:"name" bson:"name"`
    Status   string   `json:"status" bson:"status"`
    Species  string   `json:"species" bson:"species"`
    Gender   string   `json:"gender" bson:"gender"`
    Origin   struct {
        Name string `json:"name" bson:"name"`
    } `json:"origin" bson:"origin"`
    Location struct {
        Name string `json:"name" bson:"name"`
    } `json:"location" bson:"location"`
    Image   string   `json:"image" bson:"image"`
    Episode []string `json:"episode" bson:"episode"`
}

var collection *mongo.Collection

// ✅ Initialize MongoDB Connection
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
    fmt.Println("✅ Connected to MongoDB!")

    collection = client.Database("rickmorty").Collection("characters")
}

// ✅ Search Character by Name (Case-Insensitive)
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

// ✅ Main Function to Start the Server
func main() {
    initMongoDB()

    r := gin.Default()
    r.GET("/character/:name", getCharacterByName)

    fmt.Println("🚀 Server running on http://localhost:8080")
    r.Run(":8080")
}
