package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// receives the querystring from the URL using Query() method and assigns it to vl
// the query string is returned as a map so we need to take a slice to access the name hence vl['name'], we use the ok variable to check whether a name has been received this is a boolean type that will return false if no name is received, running the if block and printing the error
// otherwise we write a concatentated string injecting the name value.
// the strings.Join() method is required as we receive 'name' as a slice not a fully formed string, Join will transform the returned values of name separating them with a comma
func Hello(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	name, ok := vl["name"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}
	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}