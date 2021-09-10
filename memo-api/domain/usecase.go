package domain

import (
	"time"

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

/*TODO: Import https://github.com/T-PWK/go-fibonacci
Buscar usando deckName
O metodo vai verificar todas as cartas do deck e filtrar aquelas que:
- O score é 0
- O ReviewedAt até hoje tem que ser maior do que a tradução do score(fibonacci)
*/
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

		cardEntities = append(cardEntities, ce)
	}

	return &cardEntities, nil

}

func (u *UseCase) CheckDeckDueCards() {

}
