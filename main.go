package main

import (
	"fmt"
	"go-server/tasks"
	"net/http"
	"strings"
)

var entityHandlers = map[string]http.HandlerFunc{
	"task": tasks.TaskHandler,
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed!")
		return
	}

	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	entity := parts[0]

	if handler, ok := entityHandlers[entity]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "Hello World!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 1337")
	if err := http.ListenAndServe(":1337", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
