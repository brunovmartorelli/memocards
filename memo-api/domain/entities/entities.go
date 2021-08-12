package entities

type Card struct {
	Front string
	Back  string
	Score int
}

type Deck struct {
	Name        string
	Description string
	Cards       *[]Card
	Size        int
}
