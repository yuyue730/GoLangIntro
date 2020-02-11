package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinaryStr(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for ; n > 0; n /= 2 {
		remain := n % 2
		result = strconv.Itoa(remain) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println("Go language Loop grammar")
	fmt.Printf("Convert decimal to Binary 5=%q 13=%q 0=%q\n",
		convertToBinaryStr(5), convertToBinaryStr(13), convertToBinaryStr(0))

	fmt.Println("Print a file")
	printFile("../file.txt")
}
