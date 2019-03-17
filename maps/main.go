package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "RED",
		"green": "GREEN",
	}
	fmt.Printf("%+v", colors)

	colors2 := map[string]string{} // key, val of types strings
	fmt.Printf("%+v", colors2)     //empty map

	//manipulation
	colors["red"] = "RREEDD"
	fmt.Printf("%+v", colors)

	//deleting
	delete(colors, "red")
	fmt.Printf("%+v", colors)

	//inserting another way
	colors["yellow"] = "YELLOW"

	printMap(colors)

}

func printMap(colors map[string]string) {

	for small, capital := range colors {
		fmt.Println(small, " : ", capital)
	}

}
