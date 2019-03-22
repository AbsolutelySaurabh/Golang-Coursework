package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://golang.org",
	}

	for _, link := range links {
		// checkLink(link)
		go checkLink(link) //here we're creating a new Go routine on each iteration
	}
}

func checkLink(link string) {
	/*
		Why Go Routine ??
		The below code run on main go routine, so it has blocked it.
		 Here we're making a requets and sit around and wait for the request to respond, and then next
		 But this is a problem, we cat wait if thre are many links, so need to use Go routines
		 So we'll use additional Go routines, to avoid this problem.
	*/
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!!")
		return
	}
	fmt.Println(link, " is Up")
}
