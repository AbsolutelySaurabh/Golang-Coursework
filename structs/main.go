package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

type contactInfo struct {
	email string
	zip   int
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

	//embedding structs
	jim := person{
		firstname: "Jim",
		lastname:  "Anderson",
		contact: contactInfo{
			email: "abs@gmail.com",
			zip:   110080, // comma needed here in case of definig multiline structs
		},
	}

	fmt.Printf("%+v", jim)

}
