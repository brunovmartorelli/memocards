package domain

import (
	"reflect"
	"testing"
	"time"

	"github.com/brunovmartorelli/memo-api/domain/entities"
	"github.com/brunovmartorelli/memo-api/repository"
)

type mockExpectedCalls struct {
	get    int
	update int
	list   int
}

func TestUseCase_UpdateCardScore(t *testing.T) {

	cardRepoMock := &repository.CardRepositoryMock{
		GetCounter:    0,
		UpdateCounter: 0,
	}

	type args struct {
		deckName string
		front    string
	}
	tests := []struct {
		name    string
		u       *UseCase
		args    args
		want    int
		wantErr bool
		calls   mockExpectedCalls
	}{
		{
			name: "Given that a card score is 0, should add one point to the card's score.",
			u: &UseCase{
				cardRepository: cardRepoMock,
			},
			args: args{
				deckName: "Teste",
				front:    "Manda bala",
			},
			want:    1,
			wantErr: false,
			calls: mockExpectedCalls{
				get:    1,
				update: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.UpdateCardScore(tt.args.front, tt.args.deckName)
			if tt.calls.get != cardRepoMock.GetCounter {
				t.Errorf("mock GetCounter called = %d, want %d", cardRepoMock.GetCounter, tt.calls.get)
				return
			}
			if tt.calls.update != cardRepoMock.UpdateCounter {
				t.Errorf("mock UpdateCounter called = %d, want %d", cardRepoMock.UpdateCounter, tt.calls.update)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.UpdateCardScore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.UpdateCardScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUseCase_FilterCardsToStudy(t *testing.T) {

	cardErrorMock := &repository.CardErrorMock{
		GetCounter:    0,
		UpdateCounter: 0,
		ListCounter:   0,
	}

	cardRepoMock := &repository.CardRepositoryMock{
		GetCounter:    0,
		UpdateCounter: 0,
		ListCounter:   0,
	}

	type args struct {
		deckName string
		today    time.Time
	}
	tests := []struct {
		name    string
		u       *UseCase
		args    args
		want    *[]entities.Card
		wantErr bool
		calls   mockExpectedCalls
	}{
		{
			name: "Given a valid deck name it should return 2 cards to study and no errors",
			u: &UseCase{
				cardRepository: cardRepoMock,
			},
			args: args{
				deckName: "BODAO",
				today:    time.Date(2021, time.August, 24, 0, 0, 0, 0, time.Local),
			},
			want: &[]entities.Card{
				entities.Card{
					Front:      "Pedrao",
					Back:       "Ta aqui",
					Score:      0,
					ReviewedAt: time.Date(2021, time.August, 10, 0, 0, 0, 0, time.Local),
				},
				entities.Card{
					Front:      "Bodao",
					Back:       "Dormiu ontem fdp",
					Score:      3,
					ReviewedAt: time.Date(2021, time.August, 22, 0, 0, 0, 0, time.Local),
				},
			},
			wantErr: false,
			calls: mockExpectedCalls{
				list: 1,
			},
		},
		{
			name: "Given an invalid deck name, it should return an error",
			u: &UseCase{
				cardRepository: cardErrorMock,
			},
			args: args{
				deckName: "FADA",
				today:    time.Time{},
			},
			want:    nil,
			wantErr: true,
			calls: mockExpectedCalls{
				list: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.FilterCardsToStudy(tt.args.deckName, tt.args.today)
			if tt.calls.get != cardErrorMock.ListCounter {
				t.Errorf("mock ListCounter called = %d, want %d", cardErrorMock.ListCounter, tt.calls.list)
				return
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.FilterCardsToStudy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.FilterCardsToStudy() = %v, want %v", got, tt.want)
			}
		})
	}
}
