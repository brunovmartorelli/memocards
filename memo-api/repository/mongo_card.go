package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Front string             `json:"front"`
	Back  string             `json:"back"`
}
