package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// receiver function
func (cards deck) show() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// slices
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// join
func (cards deck) toString() string {
	temp := []string(cards)
	return strings.Join(temp, ",")
}

func (cards deck) saveToFile(filename string) error {
	bytes := []byte(cards.toString())
	return ioutil.WriteFile(filename, bytes, 0666)
}

func newDeckFromFile(filename string) deck {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	temp := string(bytes)
	return strings.Split(temp, ",")
}
