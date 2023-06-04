package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()

	assert.Equal(t, 52, len(deck), "Deck should have 52 cards")
}
