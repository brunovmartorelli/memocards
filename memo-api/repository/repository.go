package repository

import (
	"github.com/brunovmartorelli/memo-api/domain/entities"
)

type CardRepository interface {
	List(deckName string) (*[]CardSchema, error)
	GetByFront(string, string) (*CardSchema, error)
	Create(string, entities.Card) error
	Update(string, string, entities.Card) (int64, error)
	Delete(string, string) (int64, error)
}

type DeckRepository interface {
	Get(ID string) (*DeckSchema, error)
	GetByName(name string) (*DeckSchema, error)
	Create(entities.Deck) error
	Update(name string, deck entities.Deck) (int64, error)
	Delete(name string) (int64, error)
	List(cards bool) (*[]DeckSchema, error)
}
