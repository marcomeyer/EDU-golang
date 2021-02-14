package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("my-cards")
	cards = newDeckFromFile("my-cards")

	cards.shuffle()
	
	fmt.Println(cards.toString())
}

