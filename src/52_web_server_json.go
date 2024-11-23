package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Define a map to represent product data (ID, type, name, author, and nested price details)
	// to be serialized as JSON.
	response := map[string]interface{}{
		"product_id": 107,
		"type":       "book",
		"name":       "On the Origin of Species",
		"author":     "Charles Darwin",
		"price": map[string]float64{
			"regular_price":  39.99,
			"discount_price": 29.99,
		},
	}

	// Encode the map into JSON and send it in the response
	json.NewEncoder(w).Encode(response)
}

func web_server_json() {
	// Registers the handler functions for each route.
	http.HandleFunc("/about/", aboutHandler)

	// Test by visiting http://localhost:8080/about/
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
