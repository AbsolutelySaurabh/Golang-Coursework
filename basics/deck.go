package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

/*
	Writing to a File.
*/
func (d deck) saveToFile(filename string) error {

	/*
		WriteFile takes parameters as :
		1. filename
		2. []byte slice
		3. write permission which is 0666 by default
	*/
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

/*
	Reading from a file
*/
func newDeckFromFile(filename string) deck {

	bs, err := ioutil.ReadFile(filename) //bs : byte slice, need to retrieve string
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
