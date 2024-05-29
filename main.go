package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello World!")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method Not Allowed!")
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port 1337")
	if err := http.ListenAndServe(":1337", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
