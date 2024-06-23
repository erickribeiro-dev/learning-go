package main

import (
	"fmt"
	"sync"
	"time"
)

// Mutexes allow us to lock and unlock the state of a goroutine,
// so only one goroutine can access a memory at a time.
var m = sync.Mutex{}
var wg2 = sync.WaitGroup{}

var dbData = []string{"db1", "db2", "db3", "db4", "db5"}
var dbResult = []string{}

func mutexes() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg2.Add(1)
		go dbCall(i)
	}

	wg2.Wait()
	fmt.Printf("Execution time: %v\n", time.Since(t0))
	fmt.Printf("Result: %v", dbResult)
}

func dbCall(i int) {
	time.Sleep(time.Millisecond * 200)
	// Without a Mutex Lock, multiple goroutines would try to append to "result" at the same time,
	// resulting in unwanted behavior.
	m.Lock()
	dbResult = append(dbResult, dbData[i])
	// Once we're done modifying the memory, unlock the mutex so other goroutines can access it.
	m.Unlock()
	wg2.Done()
}
