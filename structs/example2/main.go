package main

import "fmt"

type Human struct{
	name string
	age int
	phone string
}

type Employee struct{
	Human	// human Hunan is also valid
	speciality string
	phone string
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

	fmt.Println("Bom work phon is ", Bob.phone)
	fmt.Println("Bob's personal phone is: ", Bob.Human.phone)


}