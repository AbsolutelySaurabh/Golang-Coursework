package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contact   contactInfo

	/*
		In this case, we need to define in this way
		contactInfo : contactInfo{
			email : "ass",
			zip : 123,
		}
	*/
	//contactInfo  --> This also works, i.e. just the type if struct and not name;
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

	//structs with receiver functions
	jim.print()
	jim.updateName("Jammy")
	jim.print()

	/*
		IN the abve statements, after updating firstname,
		it doen't changed the original

		Reason : Pass By Value

		Note : But in case of arrays, slices they get passed by reference by default.
				In this case it'll copy in memoery too but will still point to the same memoery address.
				So, in case of arrays, slices a "reference" is passed, in case of other data types value is passsed by default.
	*/

	// using pointers to avoid above
	jimPointer := &jim
	jimPointer.updateName("Samuels")
	jim.print()
}

func (pointer *person) updateName(newf string) {
	pointer.firstname = newf
}

func (p person) print() {
	fmt.Println(p)
}
