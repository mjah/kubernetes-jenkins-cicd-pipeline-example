package main

import (
	"fmt"
	"net/http"
)

// HelloWorld functions responds with Hello World and the current URL path
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! Current URL path: %s", r.URL.Path)
}

// main is the entry point which runs the Hello World server
func main() {
	// Route all requests to HelloWorld
	http.HandleFunc("/", HelloWorld)

	// Listening on port 9094
	http.ListenAndServe(":9094", nil)
}
