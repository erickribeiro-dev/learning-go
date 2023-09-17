package main

import (
	"fmt"
	"sync"
)

// Wait groups synchronizes multiple goroutines togheter
var wg = sync.WaitGroup{}
var counter = 0

func goroutines() {
	// Use the Keyword Go to run the function in a separated, green thread
	// Notice the race condition - the execution is asynchronous
	go printTask("First")
	go printTask("Second")
	go printTask("Third")

	// The main function doesn't wait for the go routine to execute,
	// that means the order of execution is unpredictable
	msg := "> Hello"
	go func() {
		fmt.Println(msg) // Can print either "Hello" or "Goodbye"
	}()
	go func(msg string) {
		fmt.Println(msg) // Can print either "Hello" or "Goodbye"
	}(msg)
	msg = "> Goodbye"

	// Using Wait Groups to remove the need of using time.Sleep on the main thread
	wg.Add(3) // Adds 3 to the Wait Group, since we're expecting 3 goroutines to finish
	go printTaskWG("Fourth")
	go printTaskWG("Fifth")
	go printTaskWG("Sixth")

	wg.Wait() // Stops waiting when the wait group reaches 0
	// time.Sleep(time.Millisecond * 100)
}

func printTask(msg string) {
	fmt.Println(msg)
}
func printTaskWG(msg string) {
	fmt.Println(msg)
	wg.Done() // Removes 1 from the Wait Group
}
