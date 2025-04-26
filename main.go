package main

import (
	"fmt"
	"time"
)

/*
We can use channels to synchronize execution across goroutines.
Here’s an example of using a blocking receive to wait for a goroutine to finish.
When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup.
*/

func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done!")
	// Send a value to notify that we’re done.
	done <- true
}

func main() {
	done := make(chan bool)

	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
}
