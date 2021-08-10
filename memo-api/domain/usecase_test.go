package domain

import "testing"

func TestUseCase_UpdateCardScore(t *testing.T) {
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
	}{
		{
			name:    "Given that a card score is 0, should add one point to the card's score.",
			u:       &UseCase{},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.UpdateCardScore(tt.args.deckName, tt.args.front)
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
