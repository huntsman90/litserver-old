package card

import (
	"math/rand"
	"time"
)

type CardDeck []Card

func InitCardDeck() CardDeck {
	var deck = make([]Card, 48)

	// Lower deck
	cardIter := 0
	for i := 0; i < 24; i++ {
		deck[cardIter].number = (i % 6) + 1
		deck[cardIter+1].number = (i % 6) + 1 + 7
		deck[cardIter].suit = i / 6
		deck[cardIter+1].suit = i / 6
		cardIter += 2
	}
	return deck
}

func (cardDeck CardDeck) Shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(deck) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}

func (cardDeck CardDeck) distributeCards(count int, groups int, shuffle bool) [][]Cards {
	if shuffle {
		cardDeck.Shuffle()
	}

	result =  [groups][count]Card

	for i:=0; i <count; i++ {
		for j:= 0 ; j<group ; j ++ {
			result[i][j] = cardDeck[(6*i)+j]
		}
	}
	return result
}
