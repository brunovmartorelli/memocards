package domain

import (
	"github.com/brunovmartorelli/memo-api/repository"
)

type UseCase struct {
	cardRepository repository.CardRepository
}

func New(cardRepository repository.CardRepository) *UseCase {
	return &UseCase{cardRepository}
}

func (u *UseCase) UpdateCardScore(deckName, front string) (int, error) {
	card, err := u.cardRepository.GetByFront(deckName, front)
	if err != nil {
		return 0, err
	}
	newScore := card.Score + 1
	return newScore, nil
}
