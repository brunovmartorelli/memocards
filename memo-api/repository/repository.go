package repository

import (
	"github.com/brunovmartorelli/memo-api/domain"
)

type CardRepository interface {
	Get(ID string) (Card, error)
	Create(domain.Card) error
	Update(domain.Card) error
	Delete(ID string) error
}

type DeckRepository interface {
	Get(ID string) (Deck, error)
	Create(domain.Deck) error
	Update(domain.Deck) error
	Delete(ID string) error
}
