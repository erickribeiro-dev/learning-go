package main

import (
	"fmt"
	"sync"
)

// A container is a map that will have multiple counters that
// can have their values changed concurrently by multiple goroutines
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Define the method inc for the type Container.
// Mutexes should not be copied, so we pass a pointer receiver instead
// inc locks the mutex state, increases the counter of map[name] by 1
// and finally unlock the mutex by the end of the inc function.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func mutexes_extra() {
	// instantiate a new container
	c1 := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// call the inc method with a name parameter that will have it's value increased n times
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c1.inc(name)
		}
		// Removes 1 from the workgroup counter
		wg.Done()
	}

	// Adds 3 to the workgroup counter, since we're calling a goroutine 3 times.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// The main thread will wait for the workgroup counter to reach 0
	wg.Wait()
	// values will show as expected
	fmt.Println(c1.counters)
}
