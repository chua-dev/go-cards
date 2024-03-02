package main

import "fmt"

func main() {

	//var card string = "Ace of Spades"
	card := "Ace of Spades" // string var init
	card = "Five of Diamonds"

	card2 := newCard()

	fmt.Println(card)
	fmt.Println(card2)

	cards := deck{newCard(), "Ace of Diamonds"}
	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	for i, v := range cards {
		fmt.Println(i, v)
	}

	cards.print()

}

func newCard() string {
	return "Five of Diamond"
}
