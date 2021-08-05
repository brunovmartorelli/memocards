package repository

import (
	"github.com/brunovmartorelli/memo-api/domain"
)

type CardRepository interface {
	Get(ID string) (*CardSchema, error)
	GetByFront(string, string) (*CardSchema, error)
	Create(string, domain.Card) error
	Update(string, string, domain.Card) (int64, error)
	Delete(string, string) (int64, error)
}

type DeckRepository interface {
	Get(ID string) (*DeckSchema, error)
	GetByName(name string) (*DeckSchema, error)
	Create(domain.Deck) error
	Update(name string, deck domain.Deck) (int64, error)
	Delete(name string) (int64, error)
	List(cards bool) (*[]DeckSchema, error)
}
