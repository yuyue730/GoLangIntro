package main

import (
	"fmt"
)

func lengthOfNonRepeatingSubstr(s string) int {
	curStart := 0
	result := 0
	lastIndexMap := make([]int, 0xffff)

	for i := range lastIndexMap {
		lastIndexMap[i] = -1
	}

	for i, ch := range []rune(s) {
		if lastIdx := lastIndexMap[ch]; lastIdx != -1 && lastIdx >= curStart {
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
