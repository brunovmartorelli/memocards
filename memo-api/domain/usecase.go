package domain

import (
	"math"
	"time"

	fib "github.com/T-PWK/go-fibonacci"
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

func (u *UseCase) FilterCardsToStudy(deckName string, today time.Time) (*[]entities.Card, error) {
	cardSchemas, err := u.cardRepository.List(deckName)
	if err != nil {
		return nil, NotFoundError{
			Err: err,
		}

	}
	var cardEntities []entities.Card

	for _, cardSchema := range *cardSchemas {

		ce := entities.Card{
			Front:      cardSchema.Front,
			Back:       cardSchema.Back,
			Score:      cardSchema.Score,
			ReviewedAt: cardSchema.ReviewedAt,
		}

		f := fib.Fibonacci(uint(ce.Score))

		daysSince := uint64(math.Floor(today.Sub(ce.ReviewedAt).Hours() / 24))

		if ce.Score == 0 || daysSince > f {
			cardEntities = append(cardEntities, ce)
		}

	}

	return &cardEntities, nil

}

func (u *UseCase) CheckDeckDueCards() {

}
