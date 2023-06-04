package main

import (
	"fmt"
	"time"
)

func main() {
	// The game logic involves players taking turns asking each other for cards of a certain rank.
	// If the asked player has cards of that rank, they give them to the asking player.
	// If not, the asking player "goes fish" and draws a card from the deck.
	// The game continues until one player has no cards left or the deck is exhausted.
	// The player with the most sets of four cards of the same rank wins.

	// Ask players for their names.
	var input string

	fmt.Println("Please enter Player 1 name: ")
	fmt.Scanln(&input)
	player_1 := Player{name: input}

	fmt.Println("Please enter Player 2 name: ")
	fmt.Scanln(&input)
	player_2 := Player{name: input}

	fmt.Println("Player 1 name is: ", player_1.name)
	fmt.Println("Player 2 name is: ", player_2.name)

	time.Sleep(1 * time.Second)

	// Create a new deck of cards
	fmt.Println("Creating new deck of cards...")
	deck := newDeck()
	// deck.print()

	time.Sleep(1 * time.Second)

	// Shuffle the deck
	fmt.Println("Shuffling deck of cards...")
	deck.shuffle()
	// deck.print()

	time.Sleep(1 * time.Second)

	// Deal each player 7 cards at the start of the game.
	fmt.Println("Dealing cards to players...")

	player_1.hand = deck.deal(7)
	fmt.Println("Player 1 hand:")
	player_1.hand.print()

	player_2.hand = deck.deal(7)
	fmt.Println("Player 2 hand:")
	player_2.hand.print()

	// Ask the player to choose a card rank to ask the other player for.
	// ..

	// If the other player has cards of that rank, they give them to the asking player.
	// ..

	// If not, the asking player "goes fish" and draws a card from the deck.
	// ..

	// The game continues until one player has no cards left or the deck is exhausted.
	// ..

	// The player with the most sets of four cards of the same rank wins.
	// ..
}
