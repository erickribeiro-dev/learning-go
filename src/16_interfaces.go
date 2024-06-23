package main

import (
	"fmt"
)

// Interfaces are a set of method signatures
// Example #1
// Defining an interface with all the methods that needs to be implemented
type Writer interface {
	Write([]byte) (int, error)
}

// Creating a new type that will implement the interface
type ConsoleWriter struct{}

// implementing the Write method in ConsoleWriter
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// The writer can do anything you implement
	n, err := fmt.Println(string(data))
	return n, err
}

// Example #2
// Any type can implement an interface (i.e., any type other than the struct type)
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

func interfaces() {
	// Creating a new ConsoleWriter instance
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Writing text with the Writer interface"))

	// Creating a new IntCounter instance
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}

	// Type switch can be used to know what types it's expected to receive
	var i interface{} = 0
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	case bool:
		fmt.Println("i is a boolean")
	default:
		fmt.Println("i is another type")
	}
}
