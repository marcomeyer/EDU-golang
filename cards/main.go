package main

import "fmt"

func main() {
	card := "Ace of Spades"
	card = newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}