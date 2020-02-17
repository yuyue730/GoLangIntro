package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello World from goroutine %d\n", i)
				runtime.Gosched() // Handout coroutine control
			}
		}(i)
	}

	time.Sleep(time.Minute)
}
