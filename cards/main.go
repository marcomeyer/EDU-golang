package main

import "fmt"

func main() {
	cards := newDeck()

	cards.saveToFile("my-cards")
	cards = newDeckFromFile("my-cards")
	
	fmt.Println(cards.toString())
}

