package main

import (
	"fmt"
	"time"
)

func channels_select() {
	ch := make(chan string)
	quit := make(chan string)
	messages := []string{"Hello", "Hi", "Hey"}
	// Sending multiple messages to the channel
	go func() {
		for _, msg := range messages {
			ch <- msg
		}
		quit <- "0"
	}()
	printMessage(ch, quit)

	// The logic after the default keyword is executed if no other case is ready
	tick := time.Tick(100 * time.Millisecond)
	end := time.After(400 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-end:
			fmt.Println("End")
			return
		default:
			fmt.Println("...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func printMessage(ch, quit chan string) {
	// Perform different actions depending on which channel the message was received from.
	for {
		select {
		case msg := <-ch:
			fmt.Println("Received message from ch:", msg)
		case <-quit:
			fmt.Println("Received message from quit")
			return
		}
	}
}
