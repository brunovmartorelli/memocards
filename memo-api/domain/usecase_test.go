package domain

import (
	"testing"

	"github.com/brunovmartorelli/memo-api/repository"
)

func TestUseCase_UpdateCardScore(t *testing.T) {

	cardRepoMock := &repository.CardRepositoryMock{
		GetCounter:    0,
		UpdateCounter: 0,
	}

	type mockExpectedCalls struct {
		get    int
		update int
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
			got, err := tt.u.UpdateCardScore(tt.args.deckName, tt.args.front)
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
