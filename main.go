package main

import (
	"fmt"
)

func main() {
	//var card string = "Ace of spades"
	// := only for defining the variable 
	//card := newCard()
	//fmt.Println(card)

	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "King of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//fmt.Println(cards)

	cards := newDeck()
	//cards.print()
	hand, remainingCards := deal(cards, 5)

	// Reciever fn eg: 
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards")

	newCards := newDeckFromFile("my_cards")
	newCards.shuffle()
	fmt.Println("Random .... ", newCards)
}
func newCard() string {
	return "Queen of Hearts" 
}