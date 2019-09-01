package main

func main() {
	// explicit variable type declaration
	// var card string = "Ace of Spades"

	// right hand variable declaration
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()

	// custom types with slices / array
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.show()

	// type conversion
	// greeting := "hola"
	// fmt.Println([]byte(greeting))

	// method declatarion and receiver function
	cards := newDeck()
	hand, _ := deal(cards, 5)
	// hand.show()
	// remaining.show()

	// save to and read from file
	hand.saveToFile("cards")
	newDeck := newDeckFromFile("cards")
	newDeck.show()
}

func newCard() string {
	return "Five of Diamonds"
}
