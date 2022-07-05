package main

import (
	"fmt"
	"strings"
)

func main() {
	findLongestSubString("pwwkew")
}

func findLongestSubString(str string) int {
	longest := 0
	seen := map[string]int{}
	start := 0

	evalString := strings.Split(str, "")
	for i, val := range evalString {
		//current letter of the string
		if _, ok := seen[val]; ok {
			start = maxVal(start, seen[val])
		}

		longest = maxVal(longest, i-start+1)
		seen[val] = i + 1
	}
	fmt.Println("---------")
	fmt.Println(longest)
	return longest
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
