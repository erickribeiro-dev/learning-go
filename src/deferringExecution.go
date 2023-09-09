package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func deferringExecution() {
	// Defer puts the function in a First In Last Out (FILO) stack
	// to be run after all the code in the scope is executed, but before returning
	defer fmt.Println("First executed")
	defer fmt.Println("Second executed")
	fmt.Println("Third executed")

	// You can defer the execution of a function if you don't need to execute it immediately
	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	//
	a := "Apple"
	// The value is decided when the function is called, not executed
	defer fmt.Println(a) // Apple
	a = "Banana"
}
