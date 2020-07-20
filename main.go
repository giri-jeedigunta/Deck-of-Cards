package main

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

	hand.print()
	remainingCards.print()
}
func newCard() string {
	return "Queen of Hearts" 
}