package tasks

import (
	"fmt"
	"net/http"
)

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Task Handler: You requested: %s", r.URL.Path)
}
