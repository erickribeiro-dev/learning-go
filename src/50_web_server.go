package main

import (
	"fmt"
	"net/http"
)

func web_server() {
	// Handler for the root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, User!")
	})

	// Handler for the /about/ route
	http.HandleFunc("/about/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "This is the about page!")
	})

	// Test by visiting http://localhost:8080 or http://localhost:8080/about/
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
