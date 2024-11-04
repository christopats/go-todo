package main

import (
	"log"
	"net/http"
)

// Our main function uses the handlefunc method to call serve file, specifying any file with the ./index.tmpl suffix
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./public"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}