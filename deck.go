package main

import (
	"fmt"
)

// Creating new type called "deck" which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ to say go that we are not going to use i and j 
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever on a fn: 
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}	
}

// func returning multiples 
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}