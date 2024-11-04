// package main

// import (
// 	"crypto/subtle"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // Constants used across the application
// // const (
// // 	CONN_HOST = "localhost"
// // 	CONN_PORT = "8080"
// // 	ADMIN_USER = "admin"
// // 	ADMIN_PASSWORD = "admin"
// // )

// // func takes ResponseWriter and Request as inputs
// // assigns the variables w and r with the values 
// func helloWorld(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }


// // func accepts .HandlerFunc to run once the auth has been accepted and a 'realm'
// 	// a realm is http terminology representing the parts of a website that require auth
// 	// a server will send a WWW-Authenticate header when auth is required

// func BasicAuth( handler http.HandlerFunc, realm string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		user, pass, ok := r.BasicAuth()
// 		// using the subtle package constanttimecompare compares the user input with the admin consts and returns 1 if the match
// 		if !ok || subtle.ConstantTimeCompare([]byte(user),
// 		[]byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass),
// 		[]byte(ADMIN_PASSWORD)) != 1 {

// 			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
// 			w.WriteHeader(401)
// 			w.Write([]byte("You are unauthorized to access the application.\n"))
// 			return
// 		}
// 		handler (w, r)
		
// 	}
// }

// func main() {
// 	http.HandleFunc("/", BasicAuth(helloWorld, "Please enter your username and password"))
// 	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
// 	if err != nil {
// 		log.Fatal("error starting http server : ", err)
// 		return
// 	}
// }