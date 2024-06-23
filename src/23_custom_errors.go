package main

import (
	"fmt"
	"time"
)

// Creating a custom error type with a timestamp and error message
type MyError struct {
	When time.Time
	What string
}

// A custom error type should implement the method Error() from the built-in "error" interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"something went wrong",
	}
}
func custom_errors() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
