package card

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Deck is a set of cards
type Deck []Card

const (
	deckSize = 48
)

// NewDeck provides a deck of 48 cards
func NewDeck() Deck {
	// The card deck always starts with 48 Cards
	// 8 Sets (4 Lower + 4 Upper) with each set containing 6 cards
	// Lower Set: 1-6, Upper Set: 8-13 ((Note 7 is skipped)
	deck := make([]Card, deckSize)

	cardIter := 0
	for i := 0; i < 24; i++ {
		// Populate entries for lower and upper deck of same suit in each iteration
		rank := (i % 6) + 1
		suit := SuitType(i / 6)

		deck[cardIter].rank = rank
		deck[cardIter].suit = suit

		deck[cardIter+1].rank = rank + 7
		deck[cardIter+1].suit = suit

		cardIter += 2
	}
	return deck
}

// DistributeCards allows to distribute given card deck among a set of Players
func (d Deck) DistributeCards(playerCount int, shuffle bool) ([][]Card, error) {
	if deckSize%playerCount != 0 {
		// return error
		return nil, errors.New("Cannot evenly distribute cards to all players")
	}
	// Shuffle is required
	if shuffle {
		d.shuffle()
	}
	cardsPerPlayer := deckSize / playerCount
	// Create Array of Card Groups to be returned
	result := make([][]Card, playerCount)
	for i := range result {
		result[i] = make([]Card, cardsPerPlayer)
	}

	// Distribute Cards from Deck in Round-Robin fashion
	for i := 0; i < cardsPerPlayer; i++ {
		for j := 0; j < playerCount; j++ {
			result[j][i] = d[(playerCount*i)+j]
		}
	}
	return result, nil
}

func (d Deck) shuffle() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}

type CardAPI struct{}

func (api CardAPI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
		return
	}

	playerCount, err := strconv.Atoi(r.FormValue("player_count"))
	if err != nil {
		log.Println("missing players in request")
		return
	}

	testDeck := NewDeck()
	cardGroup, err := testDeck.DistributeCards(playerCount, true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for i := 0; i < playerCount; i++ {
		fmt.Println("Cards for Player", i+1)
		for j := 0; j < deckSize/playerCount; j++ {
			fmt.Println(cardGroup[i][j])
		}
	}

}
