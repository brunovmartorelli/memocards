package domain

type Card struct {
	Front string
	Back  string
}

type Deck struct {
	Name        string
	Description string
	Cards       *[]Card
	Size        int
}
