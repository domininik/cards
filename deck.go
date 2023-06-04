package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	Suit string
	Rank string
}

type Deck []Card

func newDeck() Deck {
	deck := Deck{}

	suits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Suit: suit, Rank: rank})
		}
	}

	return deck
}

func (deck Deck) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func (deck *Deck) deal(handSize int) (hand Deck) {
	hand = (*deck)[:handSize]
	*deck = (*deck)[handSize:]
	return hand
}

func (deck Deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range deck {
		newPosition := r.Intn(len(deck) - 1)

		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}
