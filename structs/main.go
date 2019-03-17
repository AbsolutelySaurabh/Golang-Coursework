package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {

	alex := person{firstname: "Alex", lastname: "Johnson"}
	fmt.Println(alex)

	//updatiing struct
	var john person
	john.firstname = "John"
	john.lastname = "Players"

	fmt.Println(john)       // output : {John Players}
	fmt.Printf("%+v", john) // output : {firstname:John lastname:Players}
}
