package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("Go language String and Rune")
	s := "Yue纽约欢迎您!"
	fmt.Println("s = ", s)

	fmt.Println("1. Convert s to a slice of byte and print each item in the slice")
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println("\n")

	fmt.Println("2. Do not convert s at all and print each element.")
	for idx, ch := range s {
		fmt.Printf("(%d %X) ", idx, ch)
	}
	fmt.Println("\n")

	fmt.Println("Rune count of s = ", utf8.RuneCountInString(s))

	fmt.Println("3. Decode each rune and print it one by one")
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("(%c %d) ", ch, size)
	}
	fmt.Println("\n")

	fmt.Println("4. Convert a slice of rune")
	for idx, ch := range []rune(s) {
		fmt.Printf("(%c %d) ", ch, idx)
	}
	fmt.Println("\n")
}
