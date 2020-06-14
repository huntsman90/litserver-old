package card

// Suit defines the category of a card
type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

func (suit Suit) String() string {
	return [...]string{"Clubs", "Diamonds", "Hearts", "Spades"}[suit]
}
