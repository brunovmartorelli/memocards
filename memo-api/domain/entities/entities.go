package entities

import "time"

type Card struct {
	Front      string
	Back       string
	Score      int
	ReviewedAt time.Time
}

type Deck struct {
	Name        string
	Description string
	Cards       *[]Card
	Size        int
}
