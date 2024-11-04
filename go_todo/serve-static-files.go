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

// type Person struct {
// 	Id string
// 	Name string
// }


// func renderTemplate(w http.ResponseWriter, r *http.Request) {
// 	person := Person{Id: "1", Name: "Foo"}
// 	parsedTemplate, _ := template.ParseFiles("templates/first-template.tmpl")
// 	err := parsedTemplate.Execute(w, person)
// 	if err != nil {
// 		log.Printf("error occured while executing template or writing it's output : ", err)
// 		return
// 	}
// }

// func main() {
// 	// FileServer from http accesses all contents of the file system from root
// 	// http.Dir() enables you to specific explicitly where to locate necessary files
// 	fileServer := http.FileServer(http.Dir("static"))
// 	// Handle registers the handler that is created by the stripPrefix method
// 	// strip prefix isolates the relevant data from the URL and passes that to fileServer
// 	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
// 	http.HandleFunc("/", renderTemplate)
// 	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
// 	if err != nil {
// 		log.Fatal("error starting server : ", err)
// 		return
// 	}
// }