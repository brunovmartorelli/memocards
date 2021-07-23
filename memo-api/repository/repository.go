package repository

import (
	"github.com/brunovmartorelli/memo-api/domain"
)

type CardRepository interface {
	Get(ID string) (CardSchema, error)
	Create(domain.Card) error
	Update(domain.Card) error
	Delete(ID string) error
}

type DeckRepository interface {
	Get(ID string) (DeckSchema, error)
	Create(domain.Deck) error
	Update(domain.Deck) error
	Delete(ID string) error
}
