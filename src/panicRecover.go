package main

import (
	"fmt"
	"log"
)

func panicRecover() {
	// Panic is similar to exceptions in most other programming languages
	// Use when the an event can't be recovered

	// Example #1
	// panic: runtime error: integer divide by zero
	// a, b := 1, 0
	// result := a / b
	// fmt.Println(result)
	// fmt.Println("Finish")

	// Example #2
	// Deferred functions run even if the program panic
	// Deffered functions that close resources will run successfully even if your program panic
	// fmt.Println("Start")
	// defer fmt.Println("End")
	// panic("Something happened")

	// Example #3
	// recover() can be used to keep the program running after a panic
	fmt.Println("Start")
	panicker()
	fmt.Println("This will run after panicker() recovers")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR:", err)
			// panic(err) // This will exit the program
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking") // This line won't execute
}
