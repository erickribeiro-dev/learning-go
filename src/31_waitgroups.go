package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Wait groups synchronizes multiple goroutines togheter
// var wg sync.WaitGroup // alternative syntax
var wg = sync.WaitGroup{}

func waitgroups() {
	// Optional setting to control the maximum number of CPUs to use
	// To detect race conditions, you can run the program as: go run -race main.go
	runtime.GOMAXPROCS(4)

	// Using Wait Groups to remove the need of using time.Sleep on the main thread.
	// Adds 3 to the Wait Group counter, since we're expecting 3 goroutines to finish.
	// We can also use wg.Add(1) inside a loop, so the counter will increase by 1 each time the loop runs
	wg.Add(3)
	go printTaskWG("First")
	go printTaskWG("Second")
	go printTaskWG("Third")

	wg.Wait() // The main thread stops waiting once the wait group reaches 0
}

func printTaskWG(s string) {
	fmt.Println(s)
	wg.Done() // Removes 1 from the Wait Group
}
