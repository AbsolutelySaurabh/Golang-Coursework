package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //
	fmt.Println(r.Form) // print info on server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello saurabh!") //write daa to response
}

func login(w http.ResponseWriter, r *http.Request){

	fmt.Println("method: ", r.Method) // get request method
	if r.Method == "GET"{
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	}else{
		r.ParseForm()
		//logic part of login
		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])
	}
}

func main(){

	http.HandleFunc("/", sayhelloName) // setting router rule
	http.HandleFunc("/login", login)

	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil{
		log.Fatal("ListenAndServe: ", err)
	}
}