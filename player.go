package main

type Player struct {
	name  string
	hand  Deck
	score int
}

func (player *Player) giveCards(rank string) (cards Deck) {
	var cardsToGive Deck
	var remainingCards Deck

	for _, card := range player.hand {
		if card.Rank == rank {
			cardsToGive = append(cardsToGive, card)
		} else {
			remainingCards = append(remainingCards, card)
		}
	}

	player.hand = remainingCards

	return cardsToGive
}

func (player *Player) hasFourOfKind(rank string) bool {
	counter := 0

	for _, card := range player.hand {
		if card.Rank == rank {
			counter += 1
		}
	}

	if counter == 4 {
		return true
	} else {
		return false
	}
}

func (player *Player) removeCards(rank string) {
	var remainingCards Deck

	for _, card := range player.hand {
		if card.Rank != rank {
			remainingCards = append(remainingCards, card)
		}
	}

	player.hand = remainingCards
}
