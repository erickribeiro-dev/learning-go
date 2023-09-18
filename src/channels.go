package main

import (
	"fmt"
	"sync"
)

var wgc = sync.WaitGroup{}

func channels() {
	// Creates a channel that accepts data of the type int
	// If a second parameter is added, Go can store several messages through the channel
	ch := make(chan int, 50)
	// doneCh := make(chan struct{}, 50)

	wgc.Add(2)
	// Notice the signature - this function can only receive data
	go func(ch <-chan int) {
		// Receives data from the channel and prints it
		// msg := <-ch
		// fmt.Println(msg)

		// To process a buffer of data, you can use a for-range loop
		// for msg := range ch {
		// 	fmt.Println(msg)
		// }

		// Manually detect if the channel is closed or not by using a second parameter
		for {
			if msg, ok := <-ch; ok {
				fmt.Println(msg, ok)
			} else {
				fmt.Println(ok)
				break
			}
		}

		// The select keyword can detect which channel sent the message,
		// and process the data accordingly
		// Loop:
		// for {
		// select {
		// case entry := <-ch:
		// case <-ch:
		// fmt.Println("received message from ch", entry)
		// fmt.Println("received message from ch")
		// case <-doneCh:
		// 	fmt.Println("received message from doneCh")
		// 	break Loop
		// }
		// }

		wgc.Done()
	}(ch)

	// Notice the signature - this function can only send data
	go func(ch chan<- int) {
		// Sends values to the channel
		ch <- 1551
		ch <- 1652
		ch <- 1663

		// Close the channel to signal the receiver that no more messages will be sent,
		// avoiding deadlocks
		close(ch)
		// doneCh <- struct{}{}
		wgc.Done()
	}(ch)

	wgc.Wait()
}
