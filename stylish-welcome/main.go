package main

import (
	"log"
	"net/http"
)


func main() {
	// Handler function to render the index.tmpl template at the root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.tmpl")
	})

	// handler to access the static route containing the css files. we strip the preceding /statics/ from the file path and use the remaining file suffix to search in the public directory for the relevant file.
	http.Handle(
		"/statics/", 
		http.StripPrefix(
			"/statics/", 
			http.FileServer(http.Dir("./public")),
		),
	)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}