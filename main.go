// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// // ✅ Character struct matching MongoDB schema
// type Character struct {
// 	ID       int      `json:"id" bson:"id"`
// 	Name     string   `json:"name" bson:"name"`
// 	Status   string   `json:"status" bson:"status"`
// 	Species  string   `json:"species" bson:"species"`
// 	Gender   string   `json:"gender" bson:"gender"`
// 	Origin   struct {
// 		Name string `json:"name" bson:"name"`
// 	} `json:"origin" bson:"origin"`
// 	Location struct {
// 		Name string `json:"name" bson:"name"`
// 	} `json:"location" bson:"location"`
// 	Image   string   `json:"image" bson:"image"`
// 	Episode []string `json:"episode" bson:"episode"`
// }

// var collection *mongo.Collection

// // ✅ Initialize MongoDB Connection (CHANGE TO YOUR VM'S PUBLIC IP IF USING GCP)
// func initMongoDB() {
// 	// Use MongoDB running on VM
	
// 	mongoURI := "mongodb://34.138.200.26:27017" // ⬅ CHANGE THIS TO YOUR GCP INSTANCE IP

// 	clientOptions := options.Client().ApplyURI(mongoURI)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal("❌ MongoDB Connection Error:", err)
// 	}

// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal("❌ MongoDB Ping Failed:", err)
// 	}
// 	fmt.Println("✅ Connected to MongoDB at", mongoURI)

// 	collection = client.Database("rickmorty").Collection("characters")
// }

// // ✅ Fetch Character by Name (Fix Case-Sensitivity)
// func getCharacterByName(c *gin.Context) {
//     name := c.Param("name") // Get user input from URL parameter
//     var characters []map[string]interface{} // To store multiple results

//     // ✅ Partial Match, Case-Insensitive
//     filter := bson.M{"name": bson.M{"$regex": fmt.Sprintf(".*%s.*", name), "$options": "i"}}

//     cursor, err := collection.Find(context.TODO(), filter)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching characters"})
//         return
//     }
//     defer cursor.Close(context.TODO())

//     for cursor.Next(context.TODO()) {
//         var character map[string]interface{}
//         if err := cursor.Decode(&character); err == nil {
//             characters = append(characters, character)
//         }
//     }

//     if len(characters) == 0 {
//         c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
//         return
//     }

//     c.JSON(http.StatusOK, characters)
// }


// // ✅ Main Function to Start the Server with CORS
// func main() {
// 	initMongoDB()

// 	r := gin.Default()

// 	// Enable CORS (Allow Requests from Frontend)
// 	r.Use(cors.Default())

// 	// API Route
// 	r.GET("/character/:name", getCharacterByName)

// 	fmt.Println("🚀 Server running on http://localhost:8080")
// 	r.Run(":8080")
// }

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"backend/config"  // ✅ Change "graph-backend" to "backend"
	"backend/routes"  // ✅ Change "graph-backend" to "backend"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: No .env file found, using system environment variables.")
	}

	config.InitMongoDB()

	router := gin.Default()
	router.Use(cors.Default())

	routes.RegisterRoutes(router)

	port := ":8080"
	fmt.Println("🚀 Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))
}
