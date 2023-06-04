# Go Fish card game

The game logic involves players taking turns asking each other for cards of a certain rank.
If the asked player has cards of that rank, they give them to the asking player.
If not, the asking player "goes fish" and draws a card from the deck.
The game continues until one player has no cards left or the deck is exhausted.
The player with the most sets of four cards of the same rank wins.

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