package main

import (
	"log"
	"net/http"
)

func web_server_static_files() {
	// Serve files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))

	// Handle requests to the root ("/") by serving the index.html file
	http.Handle("/", fs)

	// Start the server on port 8080
	log.Println("Server started on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
