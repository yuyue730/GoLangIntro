package main

import (
	"fmt"
	"regexp"
)

const text = `
My email address is yuyue730@gmail.com
email1 is abc@def.org
email2 is    kk@qq.com
email3 is  ddd@abc.com.cn
`

func main() {
	fmt.Println("Go Programming Language RegExp demo")
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	fmt.Println("1. Find All String")
	match := re.FindAllString(text, -1)
	fmt.Println(match)

	fmt.Println("\n2. Find All String Submatch")
	match2 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	for _, m := range match2 {
		fmt.Println(m)
	}
}
