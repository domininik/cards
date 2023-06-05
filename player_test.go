package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGiveCards(t *testing.T) {
	player := Player{
		name: "Player 1",
		hand: Deck{
			Card{Suit: "Spades", Rank: "Ace"},
			Card{Suit: "Hearts", Rank: "Ace"},
			Card{Suit: "Diamonds", Rank: "Ace"},
			Card{Suit: "Clubs", Rank: "Ace"},
			Card{Suit: "Spades", Rank: "King"},
			Card{Suit: "Hearts", Rank: "King"},
			Card{Suit: "Diamonds", Rank: "King"},
		},
	}

	cards := player.giveCards("Ace")

	assert.Equal(t, 4, len(cards), "Player should have 4 cards")
	assert.Equal(t, 3, len(player.hand), "Player should have 3 cards")
	assert.Equal(t, "Spades", cards[0].Suit, "First card should be Spades")
	assert.Equal(t, "King", player.hand[2].Rank, "Last card should be King")
}

func TestRemoveCards(t *testing.T) {
	player := Player{
		name: "Player 1",
		hand: Deck{
			Card{Suit: "Spades", Rank: "Ace"},
			Card{Suit: "Hearts", Rank: "Ace"},
			Card{Suit: "Diamonds", Rank: "Ace"},
			Card{Suit: "Clubs", Rank: "Ace"},
			Card{Suit: "Spades", Rank: "King"},
			Card{Suit: "Hearts", Rank: "King"},
			Card{Suit: "Diamonds", Rank: "King"},
		},
	}

	player.removeCards("Ace")

	assert.Equal(t, 3, len(player.hand), "Player should have 3 cards")
	assert.Equal(t, "Spades", player.hand[0].Suit, "First card should be Spades")
	assert.Equal(t, "King", player.hand[2].Rank, "Last card should be King")
}

func TestHasFourOfKind(t *testing.T) {
	player := Player{
		name: "Player 1",
		hand: Deck{
			Card{Suit: "Spades", Rank: "Ace"},
			Card{Suit: "Hearts", Rank: "Ace"},
			Card{Suit: "Diamonds", Rank: "Ace"},
			Card{Suit: "Clubs", Rank: "Ace"},
			Card{Suit: "Spades", Rank: "King"},
			Card{Suit: "Hearts", Rank: "King"},
			Card{Suit: "Diamonds", Rank: "King"},
		},
	}

	assert.True(t, player.hasFourOfKind("Ace"), "Player should have four of a kind")
	assert.False(t, player.hasFourOfKind("King"), "Player should not have four of a kind")
}
