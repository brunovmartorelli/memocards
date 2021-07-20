package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Deck struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Cards       []Card             `json:"cards"`
	Size        int                `json:"size"`
}
