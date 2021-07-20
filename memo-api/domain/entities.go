package domain

type Card struct {
	ID    string
	Front string
	Back  string
}

type Deck struct {
	ID          string
	Name        string
	Description string
	Cards       *[]Card
	Size        int
}
