package domain

import (
	"log"

	"github.com/brunovmartorelli/memo-api/domain/entities"
	"github.com/brunovmartorelli/memo-api/repository"
)

type UseCase struct {
	cardRepository repository.CardRepository
}

func New(cardRepository repository.CardRepository) *UseCase {
	return &UseCase{cardRepository}
}

func (u *UseCase) UpdateCardScore(front, deckName string) (int, error) {
	log.Printf("usecase >> %s, %s", front, deckName)
	card, err := u.cardRepository.GetByFront(front, deckName)
	if err != nil {
		return 0, NotFoundError{
			Err: err,
		}
	}

	newScore := card.Score + 1
	card.Score = newScore

	ce := entities.Card{
		Front: card.Front,
		Back:  card.Back,
		Score: card.Score,
	}

	u.cardRepository.Update(front, deckName, ce)
	return newScore, nil
}
