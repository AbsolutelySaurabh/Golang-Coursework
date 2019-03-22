package main

/*
	//basic form

	import "fmt"
	import "net/http"


	//group form

	import(

		"fmt"
		"net/http"
	)
*/

import "fmt"

func main(){

	/*

		//basic form

		const i = 10
		const pi = 3.14
		const prefix = "Go_"

		var i int
		var pi float32
		var prefix string

		*/

		//group for
		const(
			i = 10
			pi = 3.14
			prefix = "Go_"
		)

		var(
			j = 10
			pii = 3.14
			prefx = "Go_"
		)

		fmt.Print(j, pii, prefx)

}