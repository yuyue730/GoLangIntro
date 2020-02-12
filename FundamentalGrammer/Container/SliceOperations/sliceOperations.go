package main

import "fmt"

func printSlice(slice []int) {
	fmt.Println("slice = ", slice, " len = ", len(slice), " cap = ", cap(slice))
}

func main() {
	fmt.Println("Go language Slice Operations")

	var slice []int
	for i := 0; i < 10; i++ {
		slice = append(slice, 2*i+1)
		printSlice(slice)
	}

	fmt.Println("slice = ", slice)

	// The make built-in function allocates and initializes an object of type
	// slice, map, or chan (only).
	// For example, make([]int, 0, 10) allocates an underlying array
	// of size 10 and returns a slice of length 0 and capacity 10 that is
	// backed by this underlying array.
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	fmt.Println("s2 = ", s2, "s3 = ", s3, "cap(s3) = ", cap(s3), "\n")

	copySlice := make([]int, 10)
	// The copy built-in function copies elements from a source slice into a
	// destination slice. Syntax copy(dst, src)
	copy(copySlice, slice)
	fmt.Println("copySlice = ", copySlice)

	fmt.Println("Remove 5 from copySlice = ", append(copySlice[:2], copySlice[3:]...))
	fmt.Println("Pop front from copySlice = ", copySlice[1:])
	fmt.Println("Pop back from copySlice = ", copySlice[:len(copySlice)-1])
}
