package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutexes allow us to lock and unlock the state of a goroutine,
// so only one goroutine can access a memory at a time.
var m2 = sync.RWMutex{}
var wg3 = sync.WaitGroup{}

var someData = []string{"db1", "db2", "db3", "db4", "db5"}
var someResult = []string{}

func rwmutexes() {
	t0 := time.Now()
	for i := 0; i < len(someData); i++ {
		wg3.Add(1)
		go dbCall2(i)
	}

	wg3.Wait()
	fmt.Printf("Execution time: %v\n", time.Since(t0))
	fmt.Printf("Result: %v", someResult)
}

func dbCall2(i int) {
	time.Sleep(time.Millisecond * 200)
	// Without a Mutex Lock, multiple goroutines would try to append to "result" at the same time,
	// resulting in unwanted behavior.
	saveResult(someData[i])
	showLog()
	// Once we're done modifying the memory, unlock the mutex so other goroutines can access it.
	wg3.Done()
}

func saveResult(r string) {
	m2.Lock()
	defer m2.Unlock()
	someResult = append(someResult, r)
}

func showLog() {
	// Checks if a full lock (m.lock()) exists before reading data.
	m2.RLock()         // Comment this line and the line below to check how the behavior changes
	defer m2.RUnlock() // Comment this line and the line above to check how the behavior changes
	fmt.Println(someResult)
}
