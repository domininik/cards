package main

import (
	"fmt"
	"time"
)

func main() {
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

	// The game continues until one player has no cards left or the deck is exhausted.

	for len(player_1.hand) > 0 && len(player_2.hand) > 0 && len(deck) > 0 {
		playTurn(&player_1, &player_2, &deck)
		playTurn(&player_2, &player_1, &deck)
		fmt.Println("Number of cards left in deck:", len(deck))
		time.Sleep(3 * time.Second)
	}

	// The player with the most sets of four cards of the same rank wins.
	// If both players have the same number of sets of four cards, the game is a tie.

	if player_1.score > player_2.score {
		fmt.Printf("%s wins!\n", player_1.name)
	} else if player_1.score < player_2.score {
		fmt.Printf("%s wins!\n", player_2.name)
	} else {
		fmt.Println("It's a tie!")
	}
}

func playTurn(currentPlayer, otherPlayer *Player, deck *Deck) {
	fmt.Printf("%s's turn\n", currentPlayer.name)
	fmt.Printf("Your hand: %v\n", currentPlayer.hand)

	// Ask the player to choose a card rank to ask the other player for.
	fmt.Println("Ask the other player for a rank:")
	var rank string
	fmt.Scanln(&rank)

	// If the other player has cards of that rank, they give them to the asking player.
	// If not, the asking player "goes fish" and draws a card from the deck.
	fmt.Println("Checking if other player has card of that rank...")
	time.Sleep(1 * time.Second)
	cards := otherPlayer.giveCards(rank)

	if len(cards) > 0 {
		fmt.Printf("%s had %d %s\n", otherPlayer.name, len(cards), rank)
		fmt.Printf("Transfering %d %s to %s\n", len(cards), rank, currentPlayer.name)
		currentPlayer.hand = append(currentPlayer.hand, cards...)
	} else {
		fmt.Println("Go fish :)")
		drawnCard := deck.deal(1)
		currentPlayer.hand = append(currentPlayer.hand, drawnCard...)
	}
	fmt.Printf("Your hand now: %v\n", currentPlayer.hand)

	if currentPlayer.hasFourOfKind(rank) {
		fmt.Printf("You have four %s, so you score a point!\n", rank)
		currentPlayer.removeCards(rank)
		currentPlayer.score += 1
	}
}
