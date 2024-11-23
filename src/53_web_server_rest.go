package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Retrieve a product by ID
// To test, send a POST request via curl or Firefox Network dev tool:
// curl http://localhost:8080/product/1
func getProductHandler(w http.ResponseWriter, r *http.Request) {
	// Simulating a product database with a map (in real apps, use a database)
	products := map[string]map[string]interface{}{
		"1": {
			"id":     "1",
			"name":   "Alice's Adventures in Wonderland",
			"price":  29.99,
			"author": "Lewis Carroll",
		},
		"2": {
			"id":     "2",
			"name":   "Frankenstein",
			"price":  39.99,
			"author": "Mary Shelley",
		},
	}

	// Extract product ID from the URL path (e.g., /product/1)
	// Slices the string starting at position 8 (the end of "/product/") until the end of the string.
	productID := r.URL.Path[len("/product/"):]
	product, exists := products[productID]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Simulate saving the product to a database (here we just print it)
	fmt.Printf("Retrieved existing product: %+v\n", products[productID])

	// Send product data as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Create a new product (send a JSON body with -d)
// curl -X POST http://localhost:8080/product/create/ -d '{"name": "New Book", "price": 19.99}' -H "Content-Type: application/json"
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method.", http.StatusMethodNotAllowed)
		return
	}

	var newProduct map[string]interface{}
	// Decode the incoming JSON request body into the newProduct map
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Simulate saving the product to a database (here we just print it)
	fmt.Printf("Created new product: %+v\n", newProduct)

	// Respond with the created product (normally you'd return a 201 status)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

// Update a product by ID (send a JSON body with -d)
// curl -X PUT http://localhost:8080/product/update/1 -d '{"name": "Updated Book", "price": 24.99}' -H "Content-Type: application/json"
func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow PUT method
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid HTTP method.", http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Path[len("/product/update/"):]
	var updatedProduct map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Simulate updating the product (in a real app, save to a database)
	updatedProduct["id"] = productID // Update the product with the ID from the URL
	fmt.Printf("Updated product with ID: %s\n", productID)

	// Respond with the updated product
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// Delete a product by id
// curl -X DELETE http://localhost:8080/product/delete/1
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow DELETE method
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid HTTP method", http.StatusMethodNotAllowed)
		return
	}

	productID := r.URL.Path[len("/product/delete/"):]
	// Simulate product deletion (in a real app, remove from a database)
	fmt.Printf("Deleted product with ID: %s\n", productID)

	// Respond with a success message
	w.WriteHeader(http.StatusNoContent) // 204 No Content (successful deletion)
}

func web_server_rest() {
	http.HandleFunc("/product/", getProductHandler)           // for GET
	http.HandleFunc("/product/create/", createProductHandler) // for POST
	http.HandleFunc("/product/update/", updateProductHandler) // for PUT
	http.HandleFunc("/product/delete/", deleteProductHandler) // for DELETE

	// Test by visiting http://localhost:8080/product/1
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
