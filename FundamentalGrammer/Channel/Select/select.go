package main

import (
	"fmt"
	"math/rand"
	"time"
)

// `worker` and `createWorker` functions are copied from `channel.go`.
func worker(id int, c chan int) {
	for n := range c {
		// The print speed is every second. So this will be slower than speed of
		// number pushed into Queue, so Queue size is expected to increment during
		// the life of the program.
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			// System will wait a random time between 0 and 1500 millisecond, and
			// push i into the channel. Then i will increment itself.
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	fmt.Println("Go Programming Language Usage of Select")
	var c1, c2 = generator(), generator()

	var worker = createWorker(0) // Worker variable handling values from Channel

	// Slice caching extra variables streamed out of Channel like a queue before
	// Feeding to worker
	var values []int

	// tm is a channel that will a time.Time value will be pushed into the
	// channel after 10 seconds
	tm := time.After(time.Second * 10)

	// Every second, push a time.Time value into the tick channel.
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			// Only if there is values inside values Slice, we need to activate
			// local channel receiver to handle the value from c1 and c2
			activeWorker = worker
			activeValue = values[0]
		}

		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			// For this case, if values is empty, worker() will not be executed
			// and also first element in values slice will not be removed.
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			// If the program does not print a value in 800 milliseconds, print
			// a `Timeout` on the console.
			fmt.Println("Timeout")
		case <-tick:
			fmt.Println("Queue length =", len(values))
		case <-tm:
			// The program will ends in 10 seconds
			fmt.Println("End of the program")
			return
		}
	}
}
