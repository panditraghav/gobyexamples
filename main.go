package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(time.Second)
		message <- "hello"
	}()
	fmt.Println("After func")
	/*
		By default sends and receives block until both the sender and receiver are ready.
		This property allowed us to wait at the end of our program for the "ping"
		message without having to use any other synchronization.
	*/
	fmt.Println("Message is: ", <-message)
}
