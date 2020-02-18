package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	// Naive implementation
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		// If ok is not `nil`, it means that channel is closed and we need
	// 		// to break
	// 		fmt.Println("Channel closed")
	// 		break
	// 	}
	// 	fmt.Printf("worker %d received %c\n", id, n)
	// }

	// The following implementation is even better since it will keep for going
	// until c is closed.
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
	}
}

func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

// In order to avoid defining implementing `chanDemo()` as a closure, we want to
// extract logic of 1), creating a channel and 2), having a corouting as an
// unanonymous function that takes in the integer value from channel.
func createWorker(id int) chan<- int {
	c := make(chan int)

	// We need to have a coroutine here to monitor from what is passed into
	// channels[i] and use it somehow, we also want to have an index passed in.
	go worker(id, c)
	return c
}

func bufferChannel() {
	// Buffered channel make function will take a second parameter indicating
	// size of the buffer. And it will push the value into the buffer to cache in
	// order to improve performace.
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	// close(c)
	// We can close the channel by calling `close(c)`, then the channel receiver
	// will always receive 0 if this is an integer channel or "" (empty string)
	// if a string channel
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Go language Channel")
	fmt.Println("1. Channel as first-class citizen")
	chanDemo()

	fmt.Println("\n2. Buffered Channel and Channel close")
	bufferChannel()
}
