package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// Creating new type called "deck" which is a slice of string
// This is my custom type:
type deck []string

// Creating a new deck of cards from slice of strings and for loop 
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// _ to say go that we are not going to use i and j 
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// reciever on a fn: 
// This function recieved d which is a type deck
// EG: cards.print() => cards here is of type deck and will be recieved by print fn
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}	
}

// func returning multiples ... no reciever.. passing cards as an arg
// EG: hand, remainingCards := deal(cards, 5)
func deal(d deck, handSize int) (deck, deck) {
	// :handleSize - range selectors generates subsets of the slice 
	return d[:handSize], d[handSize:]
}

// reciever on a fn: 
// util fn fo convert to string ... 
func (d deck) toString() string {
	//[]string(d) -> type conversation deck to string 
	// strings package 
	return strings.Join([]string(d), ",")
}

// Writing to file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Takes filename as input if there is error the program stops / exits 
// Reading from file 
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR : ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}

// Shuffle fn: 
func (d deck) shuffle() {
	//fmt.Println("timestamp : ", time.Now().UnixNano())
	// The below is needed to make the source unique every time the program runs 
	source := rand.NewSource(time.Now().UnixNano())
	// unique source is passed on to the rand to get new order every time this program executes
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d)-1)
		// Potision swapping in oneline !
		d[i], d[newPos] = d[newPos], d[i]
	}
}