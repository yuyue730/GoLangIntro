package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go language Map")

	map1 := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad"}

	fmt.Println("Map1 = ", map1)

	emptyMap := make(map[string]int) //We can also use `var emptyMap map[string]int`
	fmt.Println("emptyMap = ", emptyMap, "\n")

	fmt.Println("Traverse the map")
	for key, value := range map1 {
		fmt.Println("Key = ", key, "; value = ", value)
	}

	fmt.Println("\nGet key = name ", map1["name"])
	if cause, ok := map1["cause"]; ok {
		fmt.Println("map[cause] = ", cause)
	} else {
		fmt.Println("get cause rc = ", ok)
	}

	fmt.Println("\nDelete site from map1")
	_, ok1 := map1["site"]
	fmt.Println("Before deleting, can find site = ", ok1)
	delete(map1, "site")
	_, ok2 := map1["site"]
	fmt.Println("After deleting, can find site = ", ok2)
}
