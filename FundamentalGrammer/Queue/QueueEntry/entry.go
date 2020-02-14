package main

import (
	queue "GoLangIntro/FundamentalGrammer/Queue"
	"fmt"
)

func main() {
	fmt.Println("Go Language Queue Implementation")
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println("Pop 1 = ", q.Pop())
	fmt.Println("Pop 2 = ", q.Pop())
	fmt.Println("q isempty = ", q.IsEmpty())
	fmt.Println("Pop 3 = ", q.Pop())
	fmt.Println("q isempty = ", q.IsEmpty())
	q.Push("abc")
	fmt.Println("Pop abc = ", q.Pop())
}
