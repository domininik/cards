# Go Fish (in Go)

The game logic of this card game involves players taking turns and asking each other for cards of a certain rank.
If the asked player has cards of that rank, they give them to the asking player.
If not, the asking player "goes fishing" and draws a card from the deck.
The game continues until one player has no cards left or the deck is exhausted.
The player with the most sets of four cards of the same rank wins.

See https://en.wikipedia.org/wiki/Go_Fish

## Getting started

Compile files with
```
go build
```

Start the game with
```
./gofish
```

## Tests
```
go test
```