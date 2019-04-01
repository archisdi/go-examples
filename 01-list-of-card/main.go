package main

func main() {
	// var card string = "Ace of Spades"
	// card = "Five of Diamonds"

	cards := deck{newCard(), "Ace of Diamond"}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "Ace of Spade"
}
