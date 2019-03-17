package main

import (
	"fmt"
	"strings"
)

type deck []string

//this below fucntion doesnt need andy receiver, as we're creating a newDeck only, logically
func newDeck() deck {

	cards := deck{}

	//creating a new Deck
	cardSuits := []string{"Spaces", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, v := range cardValues {
		for _, s := range cardSuits {
			cards = append(cards, v+" of "+s)
		}
	}
	return cards
}

//multiple return values:
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/*
	Joinng a Slice of strings
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // this takes a slice of string as input, ans separator
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
