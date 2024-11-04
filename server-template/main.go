package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Create a struct with all the field values composing a customer
type Customer struct {
	ID int
	Name string
	Surname string
	Age int
	}	
	
// Create the template as a single string of html
// the if statement determines whether a valid id is inputted and returns customer info if so otherwise returns an error "Data not available"
// All fields within the struct are accessed with the relevant dot notation and are each checked with their own conditional
// note the importance of {{ end }} to close blocks of code.
var tplStr = `
<html>
	<h1>Customer {{ .ID }}</h1>
	{{ if .ID }}
	<p>Details:</p>
	<ul>
	{{ if .Name }}<li>Name: {{ .Name }}</li>{{ end }}
	{{ if .Surname }}<li>Surname: {{ .Surname }}</li>{{ end }}
	{{ if .Age }}<li>Age: {{ .Age }}</li>{{ end }}
	</ul>
	{{ else }}
	<p>Data not available</p>
	{{ end }}
<html>
`	
// vl receives the data inputted in the querystring and the customer struct is assigned to the cust variable
// each field of the struct performit's own conditional to check validity of data, use strconv.Atoi to convert the inputted int into string 

func Hello(w http.ResponseWriter, r *http.Request) {
	vl := r.URL.Query()
	cust := Customer{}
	id, ok := vl["id"]
	if ok {
		cust.ID, _ = strconv.Atoi(strings.Join(id, ","))
	}	
	name, ok := vl["name"]
	if ok {
		cust.Name = strings.Join(name, ",")
	}	
	surname, ok := vl["surname"]
	if ok {
		cust.Surname = strings.Join(surname, ",")
	}	
	age, ok := vl["age"]
	if ok {
		cust.Age, _ = strconv.Atoi(strings.Join(age, ""))
	}	

	// We create a template object called test which we parse our html string into and assign it to the variable tmpl
	tmpl, _ := template.New("test").Parse(tplStr)
	// run tmpl with the execute method passing the response writer and the cust struct to it
	tmpl.Execute(w, cust)
}	

// simply assign the hello handler function to our index route and starting the server
func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
