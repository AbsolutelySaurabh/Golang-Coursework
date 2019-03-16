/*
Package "main" is required to get the executable code after buiding go code.
go build name.go :  is used to compile the go code.
If we do not use "main", we'll not get the executable code.
*/
package main

import "fmt"

func main() {
	fmt.Println("Hey!!")

	/*
		Variable Decl.
	*/
	var card string = "Saurabh"
	fmt.Println(card)

	c := "Saurabh Singh"
	fmt.Println(c)
	//The below will produe error
	// c := "again"

	/*
		Functions :
	*/
	newName := newCard()
	fmt.Println(newName)

	/*
	 Arrays and Slice in go.

	 Array  :priitive ds.
	 Slice : similar to array but ore fancy.
	 example : array fixed length in go, but in case of Slice it's like ArrayList
	*/

	//	Slice :
	cards := []string{"Ace", newCard()}
	fmt.Println(cards)
	//adding a new element to slice
	cards = append(cards, "Diamonds")
	fmt.Println(cards)
	//iterating over a slice
	for i, card := range cards {
		fmt.Println(i, card)
	}

	/*
		Custom Type declaratons
	*/

	cards2 := deck{"Diamonds", "Aces", "Knights"}
	cards2 = append(cards2, "Bishop")
	cards2.print()

}

func newCard() string {
	return "Neelam Sonali"
}
