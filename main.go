package main

import "fmt"

func main() {

	//var card string = "Ace of Spades"
	card := "Ace of Spades" // string var init
	card = "Five of Diamonds"

	card2 := newCard()

	cards := []string{newCard(), "Ace of Diamonds"} // slice init
	cards = append(cards, "Six of Spades")

	for i, v := range cards {
		fmt.Println(i, v)
	}

	fmt.Println(card)
	fmt.Println(card2)
	fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamond"
}
