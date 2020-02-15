package main

import (
	"fmt"
)

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	fmt.Println("Go language Functional Programming with Adder")

	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("Add i = %d result = %d\n", i, a(i))
	}
}