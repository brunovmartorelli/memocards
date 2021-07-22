package repository

import (
	"github.com/brunovmartorelli/memo-api/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeckSchema struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Cards       []Card             `json:"cards"`
	Size        int                `json:"size"`
}

type Deck struct {
	Database   string
	Collection string
}

func (d *Deck) Get(ID string) (Deck, error) {
	return Deck{}, nil
}
func (d *Deck) Create(domain.Deck) error {
	return nil
}
func (d *Deck) Update(domain.Deck) error {
	return nil
}
func (d *Deck) Delete(ID string) error {
	return nil
}
