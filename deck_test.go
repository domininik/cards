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
