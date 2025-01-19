package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Status   string             `bson:"status" json:"status"`
	Species  string             `bson:"species" json:"species"`
	Gender   string             `bson:"gender" json:"gender"`
	Origin   struct {
		Name string `bson:"name" json:"name"`
	} `bson:"origin" json:"origin"`
	Location struct {
		Name string `bson:"name" json:"name"`
	} `bson:"location" json:"location"`
	Image   string   `bson:"image" json:"image"`
	Episode []string `bson:"episode" json:"episode"`
}
