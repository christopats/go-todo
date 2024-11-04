package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

var tasks []string
var tasksMutex sync.Mutex

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tasksMutex.Lock()
	defer tasksMutex.Unlock()

	tmpl.ExecuteTemplate(w, "index.html", tasks)

}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		r.ParseForm()
		newTask := r.FormValue("tasks")
		fmt.Println("Received task:", newTask)

		tasksMutex.Lock()
		tasks = append(tasks, newTask)
		tasksMutex.Unlock()

		tmpl.ExecuteTemplate(w, "tasks.html", tasks)
	}
}

func main() {
	fmt.Println("Templates loaded successfully")
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks", tasksHandler)

	func() error {
		server := &http.Server{Addr: ":8080", Handler: http.Handler(nil)}
		return server.ListenAndServe()
	}()
}
