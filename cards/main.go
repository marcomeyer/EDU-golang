package main

import "fmt"

func main() {
	cards := deck { "Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	
	for _, card := range cards {
		fmt.Println(card)
	}
	
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