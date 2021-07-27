package repository

import (
	"github.com/brunovmartorelli/memo-api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CardSchema struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Front string             `json:"front"`
	Back  string             `json:"back"`
}

type MongoCard struct {
	Database   string
	Collection string
	Client     *mongo.Client
}

func NewCard(c *mongo.Client) CardRepository {
	return &MongoCard{
		Client:     c,
		Database:   "memo",
		Collection: "card",
	}
}

func (d *MongoCard) Get(ID string) (CardSchema, error) {
	return CardSchema{}, nil
}
func (d *MongoCard) Create(domain.Card) error {
	return nil
}
func (d *MongoCard) Update(domain.Card) error {
	return nil
}
func (d *MongoCard) Delete(name string) (int64, error) {
	return 0, nil
}
