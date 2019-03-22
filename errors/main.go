package main

import (
	"errors"
	"fmt"
)

func main(){

	err := errors.New("This is an error!");
	if err != nil{
		fmt.Println(err)
	}
}