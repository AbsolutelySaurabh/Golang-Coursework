package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Employee struct{
	Human	// human Hunan is also valid, this is called anonmous field
	speciality string
	phone string
}

type Student struct{
	Human //anonmous field
	stream string
	phone string
}

/*
	This is basically example of method " OVERRIDING "

	means, same functions being called differently depending upon the type of receiver
 */

func (e Employee) printPhone(){
	fmt.Println("The employee is: ", e.phone)
}

func (h Human) printPhone(){
	fmt.Println("The personal phone is: ", h.phone)
}

/*
		Inheritance of methods :

		Basically, here we're using the same method is of anonmous field on diff structs using it
		Note : This happens only when it's anonmous, not in case of "h Human", is it is then we've to use Yash.h.setPhone()
 */

func (h *Human) setPhone(p string){
	h.phone = p;
}


func main(){

	Bob := Employee{
		Human{
			"Bob",
			23,
			"12345",
		},
		"Designer",
		"223-344",
	}

	Yash := Student{
		Human{
			"YASH",
			12,
			"4324",
		},
		"Science",
		"7654",
	}

	fmt.Println("Bom work phon is ", Bob.phone)
	fmt.Println("Bob's personal phone is: ", Bob.Human.phone)

	Yash.setPhone("xxxx");
	Bob.setPhone("yyyy");

	fmt.Println("Yash Human Phone: ", Yash.Human.phone)
	fmt.Println("Bob Human Phone: ", Bob.Human.phone)
}