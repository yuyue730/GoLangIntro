package main

import (
	"fmt"
)

func updateSlice(slice []int) {
	slice[0] = 100
}

func main() {
	fmt.Println("Go language Slice")

	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[:] = ", arr[:])

	s1 := arr[2:6]
	s2 := arr[:]

	updateSlice(s1)
	updateSlice(s2)
	fmt.Println("after update s1 = ", s1)
	fmt.Println("after update s2 = ", s2)
	s3 := s1[1:3]
	fmt.Println("reslice s2 to s3 = s1[1:3] = ", s3, "\n")

	fmt.Println("after update arr = ", arr)
	fmt.Println("Print slice members for s1, s2 and s3")
	// Cap official documentation:
	// The cap built-in function returns the capacity of v, according to its type:
	//	Slice: the maximum length the slice can reach when resliced;
	//	if v is nil, cap(v) is zero.
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))

	// The append built-in function appends elements to the end of a slice. If
	// it has sufficient capacity, the destination is resliced to accommodate the
	// new elements. If it does not, a new underlying array will be allocated.
	// Append returns the updated slice. It is therefore necessary to store the
	// result of append, often in the variable holding the slice itself.
	fmt.Println("Append to s1, 10, 11, 12")
	append1 := append(s1, 10)
	append2 := append(append1, 11)
	append3 := append(append2, 12)
	fmt.Println("append1 = ", append1)
	fmt.Println("append2 = ", append2)
	fmt.Println("append3 = ", append3)
	fmt.Println("arr = ", arr)
}
