package main

import (
	"fmt"
	"net/http"
)

// Basic handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, User!")
	// Alternative writer:
	// w.Write([]byte("Hello, User!"))
}

// Handler with query parameters
// Example: http://localhost:8080/product?name=White+Shirt&size=40
func productHandler(w http.ResponseWriter, r *http.Request) {
	// Set Content-Type to text/html
	w.Header().Set("Content-Type", "text/html")

	// Get query parameters.
	name := r.URL.Query().Get("name")
	size := r.URL.Query().Get("size")

	// Send an HTML message.
	if name == "" || size == "" {
		fmt.Fprintf(w, "<h1>Product not found.</h1><p>Please send the parameters: name and size</p>")
		return
	}

	fmt.Fprintf(w, "<p>You are viewing the product: <b>%s</b> of size: <b>%s</b>.</p>", name, size)
}

func web_server_query() {
	// Registers the handler functions for each route.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/product/", productHandler)

	// Test by visiting http://localhost:8080/product/?name=White+Shirt&size=40
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
