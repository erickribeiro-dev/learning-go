package main

import (
	"fmt"
	"net/http"
)

func web_server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
		fmt.Println("Got a new request.")
	})

	// Start the server on port 8080
	// To test, open your browser to http://localhost:8080 or 127.0.0.1:8080
	http.ListenAndServe(":8080", nil)
}
