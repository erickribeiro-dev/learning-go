package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func channels_extra() {
	// Example 1: divide the work into 2 goroutines
	dataCh := make(chan int)
	intSliceOne := []int{1, 2, 3, 4, 5}
	intSliceTwo := []int{6, 7, 8, 9, 10}
	go sum(intSliceOne, dataCh)
	go sum(intSliceTwo, dataCh)
	//Receive the values from both the channel and print them
	x, y := <-dataCh, <-dataCh
	fmt.Printf("x = %d; y = %d; x+y = %d\n", x, y, x+y)

	// Example 2:
	go func() {
		// add a wait group to track when the goroutines are done so the code can continue (after wg.Wait())
		wg := sync.WaitGroup{}
		for i := 0; i < 100; i++ {
			// Instead of waiting for the processing in the main thread,
			// wrap the processing in goroutines for each iteration of the loop
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataCh <- getRandomNum()
			}()
		}
		wg.Wait()
		// Close the channel
		close(dataCh)
	}()
	for v := range dataCh {
		fmt.Printf("v = %d\n", v)
	}

	// Example 3: Calculate Fibonacci using channels
	fibChannel := make(chan int, 20)
	go fibonacci(cap(fibChannel), fibChannel)
	for i := range fibChannel {
		fmt.Println(i)
	}

}

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	ch <- sum
}

func getRandomNum() int {
	// sleep for 1 second to simulate heavy processing, so we can observe and compare the performance
	time.Sleep(time.Second)
	return rand.Intn(100)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
