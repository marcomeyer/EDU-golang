package main

import "fmt"

func main() {
	card := "Ace of Spades"
	card = newCard()

	fmt.Println(card)
	printState()
	loopWithUnusedIndex()
}

func newCard() string {
	return "Five of Diamonds"
}

func loopWithUnusedIndex() {
	
	cards := [2]string {"A", "B"}
	for _ ,card := range cards {
		fmt.Println(card)
	}
}