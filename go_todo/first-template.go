// package main

// import (
// 	"html/template"
// 	"log"
// 	"net/http"
// )

// const (
// 	CONN_HOST = "localhost"
// 	CONN_PORT = "8080"
// )

// // Created a data structure containing Id and Name which are referenced in the .tmpl file
// type Person struct {
// 	Id string
// 	Name string
// }

// // function to render the template from the templates/ dir
// func renderTemplate(w http.ResponseWriter, r *http.Request) {
// 	// var person is initilised and receives data
// 	person := Person{Id: "1", Name: "Foo"}
// 	// calling the ParseFiles method from html/templates, this creates a new template at the variable parsedTemplate
// 	parsedTemplate, _ := template.ParseFiles("templates/first-template.tmpl")
// 	// Now we execute the newly created template passing in the person variable and responsewriter as args. this generates the html and passes it to the http response stream
// 	err := parsedTemplate.Execute(w, person)
// 	if err != nil {
// 		log.Printf("error occured while executing template or writing it's output : ", err)
// 		return
// 	}
// }

// func main() {
// 	// here our index route calls the renderTemplate function
// 	http.HandleFunc("/", renderTemplate)
// 	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
// 	if err != nil {
// 		log.Fatal("Could not connect to server : ", err)
// 		return
// 	}
// }