package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a struct containing all the data we want to display on our pages
type PageWithCounter struct {
	Counter int `json:"counter"`
	Heading string `json:"heading"`
	Content string `json:"content"`
}

// Create a handler function that serves the content to all routes, dynamically inputting it into the relevant field using struct dot notation
// Also incrementing the h.Counter each time the function is called allows a page views counter
// using the pointer of the PageWithCounter struct enables the data stored there to be altered with each call as opposed to calling the struct intself which would make a copy, resulting in the views counter displaying 1 each time page refreshed
func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter++
	msg := fmt.Sprintf("<h1>%s</h1>\n<p>%s</p>\n<p>Views: %d</p>", h.Heading, h.Content, h.Counter)
	w.Write([]byte(msg))
}

// Create instances of our PageWithCounter struct naming them with appropriate references to their purpose/route
func main() {
	hello := PageWithCounter{
		Heading: "Hello World",
		Content: "This is the main page"}
	cha1 := PageWithCounter{
		Heading: "Chapter 1",
		Content: "This is the first Chapter"}
	cha2 := PageWithCounter{
		Heading: "Chapter 2",
		Content: "This is the second Chapter"}

// Handle each route, using the pointer location so as to enable the incrementing of the views counter 
	http.Handle("/", &hello)
	http.Handle("/chapter1", &cha1)
	http.Handle("/chapter2", &cha2)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
