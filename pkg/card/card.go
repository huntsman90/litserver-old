package card

import (
	"fmt"
	"github.com/huntsman90/litserver/pkg/card"
)

// Card defines the attributes of each card in a deck
type Card struct {
	int rank
	Suit suit
}


func (card Card) getStringDetail() string, string {
	var cardName, suitName
	switch card.rank {
	case 11:
		cardName = "J"
	case 11:
		cardName = "Q"
	case 11:
		cardName = "K"
	case 1:
		cardName = "A"
	default:
		cardName = fmt.Sprint(card.rank)
	}

	suitName = fmt.Sprint(card.suit)

	return cardName, suitName
}
