package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	assert.Equal(t, 52, len(deck), "Deck should have 52 cards")

	assert.Equal(t, "Spades", deck[0].Suit, "First card should be Spades")
	assert.Equal(t, "Ace", deck[0].Rank, "First card should be Ace")

	assert.Equal(t, "Clubs", deck[51].Suit, "Last card should be Clubs")
	assert.Equal(t, "King", deck[51].Rank, "Last card should be King")
}

func TestDeal(t *testing.T) {
	deck := newDeck()

	hand := deck.deal(7)

	assert.Equal(t, 7, len(hand), "Hand should have 7 cards")
	assert.Equal(t, 45, len(deck), "Deck should have 45 cards")
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	anotherDeck := newDeck()

	assert.Equal(t, deck, anotherDeck, "Decks should be equal")

	deck.shuffle()

	assert.NotEqual(t, deck, anotherDeck, "Decks should not be equal")
}
