package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()
	hand, remaining := deal(cards, 5)

	cards.show()
	hand.show()
	remaining.show()
}

func newCard() string {
	return "Five of Diamonds"
}
