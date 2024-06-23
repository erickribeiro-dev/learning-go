package main

import (
	"fmt"
	"time"
)

// Goroutines are a concurrent, lightweight threads of execution
func goroutines() {
	someData := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// functions with the go keyword will run using a goroutine
	// Let's divide our processing data in 2, so it can be processed by concurrently by 2 goroutines
	go heavyProcessingFunc(someData[:len(someData)/2]) // 0th index to "size/2"th index
	go heavyProcessingFunc(someData[len(someData)/2:]) // "size/2"th index to last index

	// Goroutines can also be used in anonymous functions.
	// The execution is unpredictable, the surrounding function (goroutines()) doesn't wait for goroutines to finish.
	msg := "Hello"
	go func() {
		fmt.Println(msg) // Will print either Hello or Goodbye
		fmt.Println("This anonymous function was executed by a goroutine")
	}()
	msg = "Goodbye"

	// The main thread will not wait for the goroutines to finish processing to exit,
	// so let's use the time.Sleep() function for now
	time.Sleep(time.Second * 3)
}

func heavyProcessingFunc(i []int) {
	for _, v := range i {
		// use time.Sleep() here to simulate a long processing
		// Notice how both goroutines run in concurrently
		time.Sleep(time.Millisecond * 500)
		fmt.Println(v, v*v)
	}
}
