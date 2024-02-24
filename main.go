package main

import "fmt"

func main() {

	//var card string = "Ace of Spades"
	card := "Ace of Spades" // initialization
	card = "Five of Diamonds"

	card2 := newCard()

	fmt.Println(card)
	fmt.Println(card2)
}

func newCard() string {
	return "Five of Diamond"
}
