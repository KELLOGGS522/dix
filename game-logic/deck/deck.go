package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Rank string `json:"rank"`
	Suit string `json:"suit"`
}

func RankSuit(c Card) string {
	return c.Rank + "-" + c.Suit
}

type Deck []Card

var Ranks = []string{
	"Ace",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
}

var Suits = []string{
	"Spade",
	"Diamond",
	"Club",
	"Heart",
}

// Create new deck of cards for the game
func New() (deck Deck) {
	// loop over each rank and suit appending to the deck
	for i := 0; i < len(Ranks); i++ {
		for j := 0; j < len(Suits); j++ {
			card := Card{
				Rank: Ranks[i],
				Suit: Suits[j],
			}
			deck = append(deck, card)
		}
	}
	return
}

var shuffleRand = rand.New(rand.NewSource(time.Now().Unix()))

// Shuffle the deck
func Shuffle(deck Deck) Deck {
	ret := make([]Card, len(deck))
	perm := shuffleRand.Perm(len(deck))
	for i, j := range perm {
		ret[i] = deck[j]
	}
	return ret
}

// Find takes a slice of Card and looks for a Card in it. If found it will
// return it's index, otherwise it will return -1 and a bool of false.
// https://golangcode.com/check-if-element-exists-in-slice/
func Find(deck []Card, c Card) (int, bool) {
	for i := range deck {
		if RankSuit(c) == RankSuit(deck[i]) {
			return i, true
		}
	}
	return -1, false
}

// RemoveCard removes a Card from a deck with immutability
// a new deck is returned and the one passed as input is left unchanged
// https://stackoverflow.com/a/59205977/5699993
func RemoveCard(deck []Card, c Card) (newDeck []Card) {
	i, found := Find(deck, c)
	if !found {
		fmt.Printf("Panicking! The card %v is not in the deck %v\n", c, deck)
		panic(fmt.Sprintf("Panicking! The card %v is not in the deck %v\n", c, deck))
	}
	for j := range deck {
		if j != i {
			newDeck = append(newDeck, deck[j])
		}
	}
	return
}

// RemoveCards removes a slice of Cards from a deck
func RemoveCards(deck []Card, cardsToRemove []Card) (newDeck []Card) {
	// copy for immutability
	newDeck = make([]Card, len(deck))
	copy(newDeck, deck)
	for _, c := range cardsToRemove {
		newDeck = RemoveCard(newDeck, c)
	}
	return
}
