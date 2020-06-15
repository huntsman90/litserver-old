package card

import (
	"fmt"
)

// Card struct holds the definition of a single card
type Card struct {
	rank int //number/alphabet on the card
	suit SuitType
}

// Override the Stringer method to print card details in human understandable format
func (c Card) String() string {
	var cardName string
	switch c.rank {
	case 11:
		cardName = "J"
	case 12:
		cardName = "Q"
	case 13:
		cardName = "K"
	case 1:
		cardName = "A"
	default:
		cardName = fmt.Sprint(c.rank)
	}

	// Ex: Return "A-Hearts" or "10-Spades"
	return cardName + "-" + fmt.Sprint(c.suit)
}

// Rank returns the rank value of a card
func (c Card) Rank() int {
	return c.rank
}

// Suit returns suit value of the card
func (c Card) Suit() SuitType {
	return c.suit
}
