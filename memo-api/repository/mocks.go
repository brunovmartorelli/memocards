package repository

import (
	"fmt"
	"time"

	"github.com/brunovmartorelli/memo-api/domain/entities"
)

type CardRepositoryMock struct {
	GetCounter    int
	UpdateCounter int
	ListCounter   int
}

func (c *CardRepositoryMock) List(deckName string) (*[]CardSchema, error) {
	c.ListCounter++
	cards := &[]CardSchema{
		CardSchema{
			Front:      "Pedrao",
			Back:       "Ta aqui",
			Score:      0,
			ReviewedAt: time.Date(2021, time.August, 10, 0, 0, 0, 0, time.Local),
		},
		CardSchema{
			Front:      "Bodao",
			Back:       "Dormiu ontem fdp",
			Score:      3,
			ReviewedAt: time.Date(2021, time.August, 22, 0, 0, 0, 0, time.Local),
		},
		CardSchema{
			Front:      "Seila",
			Back:       "Mano",
			Score:      5,
			ReviewedAt: time.Date(2021, time.August, 24, 0, 0, 0, 0, time.Local),
		},
		CardSchema{
			Front:      "Irineu",
			Back:       "VC N SB NI EU",
			Score:      7,
			ReviewedAt: time.Date(2021, time.August, 24, 0, 0, 0, 0, time.Local),
		},
		CardSchema{
			Front:      "Bathtub",
			Back:       "Twitch",
			Score:      10,
			ReviewedAt: time.Date(2021, time.August, 24, 0, 0, 0, 0, time.Local),
		},
	}

	return cards, nil
}
func (c *CardRepositoryMock) GetByFront(string, string) (*CardSchema, error) {
	c.GetCounter++
	return &CardSchema{
		Front: "Manda bala",
		Back:  "Okay",
		Score: 0,
	}, nil
}
func (c *CardRepositoryMock) Create(string, entities.Card) error {
	return nil
}
func (c *CardRepositoryMock) Update(string, string, entities.Card) (int64, error) {
	c.UpdateCounter++
	return 0, nil
}
func (c *CardRepositoryMock) Delete(string, string) (int64, error) {
	return 0, nil
}

type CardErrorMock struct {
	GetCounter    int
	UpdateCounter int
	ListCounter   int
}

func (c *CardErrorMock) List(deckName string) (*[]CardSchema, error) {
	c.ListCounter++
	return nil, fmt.Errorf("failed to list deck")
}
func (c *CardErrorMock) GetByFront(front, deckName string) (*CardSchema, error) {
	c.GetCounter++
	return nil, fmt.Errorf("deck %s or card %s not found", deckName, front)
}
func (c *CardErrorMock) Create(string, entities.Card) error {
	return nil
}
func (c *CardErrorMock) Update(string, string, entities.Card) (int64, error) {
	c.UpdateCounter++
	return 0, nil
}
func (c *CardErrorMock) Delete(string, string) (int64, error) {
	return 0, nil
}
