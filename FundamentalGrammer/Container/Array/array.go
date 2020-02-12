package main

import (
	"fmt"
)

func printArray(arr [5]int) {
	for i, v := range arr {
		fmt.Println("Index = ", i, " value = ", v)
	}
}

func main() {
	fmt.Println("Go language Array")
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println("arr1 = ", arr1)
	fmt.Println("arr2 = ", arr2)
	fmt.Println("arr3 = ", arr3)

	var grid [4][5]int
	fmt.Println("grid = ", grid)

	fmt.Println("Print Array items for arr2")
	printArray(arr3)
}
