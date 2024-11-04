package main

import (
	"log"
	"net/http"
)

// create handler functions for 2 routes both displaying the same inital part of a message

// Use a function with the parameters of the existing handler functions to be called when the route is access which in turn after it's action is performed it will call the function contained within it's argument

// This is the middleware method function, it accepts a handler func type as args and names it next, also returning a handler func method
// The function declared within the middleware writes the hello there message, then calls serve http on the handler function pass in as args
func Hello(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "Hello there, "
		w.Write([]byte(msg))
		next.ServeHTTP(w, r)
	}
}

// function declarations for the action at each specific route, these are then passed to Hello which acts as middleware, performing it's own task before then calling these functions
func Function1(w http.ResponseWriter, r *http.Request) {
		msg := "this is function 1"
		w.Write([]byte(msg))
	}

func Function2(w http.ResponseWriter, r *http.Request) {
		msg := "this is function 2"
		w.Write([]byte(msg))
	}

// We assign the routes calling Hello each time and passing functions 1 or 2 in respectively 
// activate the 8080 port to listen and serve
func main() {

	http.HandleFunc("/hello1", Hello(Function1))
	http.HandleFunc("/hello2", Hello(Function2))
	log.Fatal(http.ListenAndServe(":8080", nil))
}