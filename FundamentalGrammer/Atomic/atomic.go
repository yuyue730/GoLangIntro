package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment() {
	fmt.Println("Safe Increment")
	func() {
		// Lock the resource using mutex before accessing critical data, value. Defer
		// the unlock until right before the method is done. Using executing an 
		// ananonymous function to lock a block inside a function.
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
	
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("a =", a.value)
}