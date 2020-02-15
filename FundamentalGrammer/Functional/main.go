package main

import (
	fib "GoLangIntro/FundamentalGrammer/Functional/Fib"
	"bufio"
	"fmt"
	"io"
	"strings"
)

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 2000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("Go language Functional Programming with Fibonacci")
	var f intGen = fib.Fibonacci()
	fmt.Println("Print all sum less than 5000")
	printFileContents(f)
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("i = %d, sum = %d\n", i, f())
	// }
}
