package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Invalid score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "file.txt"
	fmt.Println("Go language If and Switch branch grammar")
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Print file content = [%s]\n", contents)
	}

	fmt.Printf("0=%q, 59=%q, 60=%q, 82=%q, 99=%q, 100=%q\n",
		grade(0), grade(59), grade(60), grade(82), grade(99), grade(100))
	fmt.Println("Expect to crash with -3")
	fmt.Printf("-3=%s\n", grade(-3))
}
