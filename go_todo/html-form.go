// package main

// import (
// 	"fmt"
// 	"html/template"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/schema"
// )

// const (
// 	CONN_HOST = "localhost"
// 	CONN_PORT = "8080"
// )

// type User struct {
// 	Username string
// 	Password string
// }


// func readForm(r *http.Request) *User {
// 	r.ParseForm()
// 	user := new(User)
// 	decoder := schema.NewDecoder()
// 	decodeErr := decoder.Decode(user, r.PostForm)
// 	if decodeErr != nil {
// 		log.Printf("error maooing parsed form data to struct : ", decodeErr)}
// 		return user	
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		parsedTemplate, _ := template.ParseFiles("templates/html-form.tmpl")
// 		parsedTemplate.Execute(w, nil)
// 	} else {
// 		user := readForm(r)
// 		fmt.Fprintf(w, "Hello "+user.Username+"!")
// 	}
// }

// func main() {
// 	fileServer := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fileServer))
// 	http.HandleFunc("/", login)
// 	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
// 	if err != nil {

// 		log.Fatal("error starting server : ", err)
// 		return
// 	}
// }

