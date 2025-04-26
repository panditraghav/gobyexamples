package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Go’s select lets you wait on multiple channel operations.
		Combining goroutines and channels with select is a powerful feature of Go.
	*/
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Second"
	}()

	// We’ll use select to await both of these values simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("recieved: ", m1)
		case m2 := <-c2:
			fmt.Println("recieved: ", m2)
		}
	}
}
