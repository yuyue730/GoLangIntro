package main

import (
	"fmt"
	"sync"
)

// func doWork(id int, c chan int, done chan bool) {
// 	// The following implementation is even better since it will keep for going
// 	// until c is closed.
// 	for n := range c {
// 		fmt.Printf("worker %d received %c\n", id, n)
// 		done <- true
// 	}
// }

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}

// worker holds two channels, `in` is the integer value `createWorker` tries to
// pass to channel receiver to `doWork`. Done is the boolean value `doWork` will
// set and pass out after work is completed
// type worker struct {
// 	in chan int
// 	done chan bool
// }

// Instead of using boolean channel done, define a sync.WaitGroup pointer or a
// done function which is `wg.Done()` for the same purpose
type worker struct {
	in chan int
	// wg *sync.WaitGroup
	done func()
}

// In order to avoid defining implementing `chanDemo()` as a closure, we want to
// extract logic of 1), creating a channel and 2), having a corouting as an
// unanonymous function that takes in the integer value from channel.
func createWorker(id int, wg *sync.WaitGroup) worker {
	// w := worker{
	// in: make(chan int),
	// done: make(chan bool),
	// }

	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() },
		// wg: wg,
	}

	// We need to have a coroutine here to monitor from what is passed into
	// channels[i] and use it somehow, we also want to have an index passed in.
	// go doWork(id, w.in, w.done)

	// Pass wg instead for new way of implementation
	go doWork(id, w)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	// Wait for each work to be done (true to be pushed into `work.done` channel)
	// Can be replaced by WaitGroup
	// for _, worker := range workers {
	// 	<-worker.done
	// }

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// Can be replaced by WaitGroup
	// for _, worker := range workers {
	// 	<-worker.done
	// }

	wg.Wait() // Wait for all WorkGroup to call `done()`
}

func main() {
	fmt.Println("Go Programming Language add done channel/WaitGroup for Coroutine",
		"to inform task has been done")
	chanDemo()
}
