package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayTurn(t *testing.T) {
	assert := assert.New(t)

	player_1 := Player{
		name: "Player 1",
		hand: Deck{
			Card{Suit: "Spades", Rank: "Ace"},
			Card{Suit: "Diamonds", Rank: "King"},
		},
	}

	player_2 := Player{
		name: "Player 2",
		hand: Deck{
			Card{Suit: "Spades", Rank: "6"},
			Card{Suit: "Clubs", Rank: "10"},
		},
	}

	deck := Deck{
		Card{Suit: "Spades", Rank: "8"},
		Card{Suit: "Hearts", Rank: "2"},
	}

	playTurn(&player_1, &player_2, &deck)

	assert.Equal(3, len(player_1.hand), "Player 1 should have 3 cards")
	assert.Equal(true, len(player_2.hand) >= 1, "Player 2 should have at least 1 card")
	assert.Equal(true, len(deck) <= 2, "Deck should have no more than 2 cards")
}
