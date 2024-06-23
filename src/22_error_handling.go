package main

import (
	"errors"
	"fmt"
	"log"
)

func error_handling() {
	defer say_goodbye()

	res, err := division_error(3, 0)
	fmt.Println(res, err)

	res, err = division_recover(3, 0)
	fmt.Println(res, err)

	// force_panic()
}

func division_error(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

// Recover from a panic
func division_recover(x, y int) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recovered from error:", err)
		}
	}()
	return x / y, nil
}

// The panic() function can be used to force the entire program to exit with a status.
// It is used to abort execution if we don't want to handle an error value
func force_panic() {
	panic("the program has panicked")
}

// Deferred functions run even if the program panic.
// Deffered functions that close resources (such as files or server connections)
// will run successfully even if your program panic
func say_goodbye() {
	fmt.Println("Goodbye!")
}
