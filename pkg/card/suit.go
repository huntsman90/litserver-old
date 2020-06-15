package card

// SuitType defines the category of a card
type SuitType int

const (
	Clubs SuitType = iota
	Diamonds
	Hearts
	Spades
)

func (s SuitType) String() string {
	return [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}[s]
}
