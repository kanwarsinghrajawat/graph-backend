package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"backend/config"
	"backend/models"

	"go.mongodb.org/mongo-driver/bson"
)

const apiURL = "https://rickandmortyapi.com/api/character/?page="

func fetchCharacters() []models.Character {
	var allCharacters []models.Character

	for page := 1; page <= 42; page++ {
		url := fmt.Sprintf("%s%d", apiURL, page)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to fetch page %d: %v", page, err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed to read response from page %d: %v", page, err)
		}

		var result struct {
			Results []models.Character `json:"results"`
		}
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Failed to parse JSON from page %d: %v", page, err)
		}

		if len(result.Results) == 0 {
			log.Fatalf(" Page %d returned 0 characters. Check API response.", page)
		}

		log.Printf("✅ Fetched %d characters from page %d", len(result.Results), page)
		allCharacters = append(allCharacters, result.Results...)
	}

	log.Printf(" Total fetched characters: %d", len(allCharacters))
	return allCharacters
}

func insertCharactersIntoDB(characters []models.Character) {
	collection := config.DB.Collection("characters")

	for _, character := range characters {
		filter := bson.M{"id": character.ID}
		count, _ := collection.CountDocuments(context.TODO(), filter)

		if count == 0 { // Only insert if not already in DB
			_, err := collection.InsertOne(context.TODO(), character)
			if err != nil {
				log.Printf("Failed to insert character %s: %v", character.Name, err)
			} else {
				log.Printf("Inserted: %s", character.Name)
			}
		} else {
			log.Printf("Skipping (already exists): %s", character.Name)
		}
	}
}

func main() {
	fmt.Println("Initializing MongoDB connection...")
	config.InitMongoDB()

	fmt.Println("Fetching characters from API...")
	characters := fetchCharacters()
	fmt.Printf("Fetched %d characters\n", len(characters))

	fmt.Println("Inserting characters into MongoDB...")
	insertCharactersIntoDB(characters)

	fmt.Println("Data seeding complete!")
}
