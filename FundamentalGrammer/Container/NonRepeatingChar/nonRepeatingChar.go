package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubstr(s string) int {
	curStart := 0
	result := 0
	lastIndexMap := make(map[byte]int)

	for i, ch := range []byte(s) {
		if lastIdx, rc := lastIndexMap[ch]; rc && lastIdx >= curStart {
			curStart = lastIdx + 1
		}
		if i-curStart+1 > result {
			result = i - curStart + 1
		}
		lastIndexMap[ch] = i
	}

	return result
}

func main() {
	fmt.Println("Go language Non Repeating Character Leetcode problem")
	fmt.Println("pwwkew Longest non repeating char length = ",
		lengthOfNonRepeatingSubstr("pwwkew"))
}
