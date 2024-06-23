package main

import (
	"fmt"
)

// Channels can be used to send and receive data between goroutines
func channels() {
	// Declare a new channel that can work with int values
	myChannel := make(chan int)
	// Send a value into a channel with the syntax: channel <- value
	// The value will be stored in the channel until it's used once.
	// Note: sending values to a channel synchronously will result in deadlock,
	// unless the channel is buffered. Therefore the data should be sent and received simultaneously.
	go func() {
		myChannel <- 5
	}()
	// Receive a value from a channel with the syntax: variable := <- channel
	// Trying to receive values more than the amount held by the channel will result in deadlock.
	myReceiver := <-myChannel
	fmt.Println(myReceiver)

	// Buffered channels can store a limited amount of values to be used whenever it's needed
	myBufferedChannel := make(chan string, 2)
	myBufferedChannel <- "Hello"
	myBufferedChannel <- "Channel"
	strReceiver := <-myBufferedChannel // Receives Hello
	fmt.Println(strReceiver)
	strReceiver = <-myBufferedChannel // Receives Channel
	fmt.Println(strReceiver)

	// Sending and receiving multiple values
	secondChan := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			secondChan <- i
		}
		// Close the channel so there's no deadlock
		close(secondChan)
	}()
	// Get values from the channel
	for v := range secondChan {
		fmt.Println(v)
	}
}
