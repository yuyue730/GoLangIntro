package main

import (
	fib "GoLangIntro/FundamentalGrammer/Functional/Fib"
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("Defer print %d\n", i)
	}
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if panicError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(panicError.Op, panicError.Path, panicError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
	defer writer.Flush()
}

func main() {
	fmt.Println("Go language Defer example")
	fmt.Println(("1. Try to show defer calling as a Stack."))
	tryDefer()
	fmt.Println()

	fmt.Println("2. Print Fibonacci to file")
	writeFile("fib.txt")
}
