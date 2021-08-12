package repository

import "github.com/brunovmartorelli/memo-api/domain/entities"

type CardRepositoryMock struct {
	GetCounter    int
	UpdateCounter int
}

func (c *CardRepositoryMock) List(deckName string) (*[]CardSchema, error) {
	return nil, nil
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
