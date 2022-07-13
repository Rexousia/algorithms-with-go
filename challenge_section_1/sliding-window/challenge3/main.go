package main

import (
	"fmt"
	"strings"
)

func main() {
	findLongestSubString("pwwkew")
	findLongestSubString("kodak")
}

func findLongestSubString(str string) int {
	longest := 0
	seen := map[string]int{}
	start := 0

	evalString := strings.Split(str, "")
	for i, val := range evalString {
		//current letter of the string
		if _, ok := seen[val]; ok {
			start = maxVal(longest, seen[val])
			fmt.Println("---------")
			fmt.Println("Inside of the condition where the value is found true")
			fmt.Println("Start:", start)
			fmt.Println("Seen[val]:", seen[val])
			fmt.Println("---------")
		}

		//substracting the current start from the current iteration and adding 1
		//evaluating to see which is longest
		longest = maxVal(longest, i-start+1)
		fmt.Println("---------")
		fmt.Println("Inside for loop outside of conditon")
		fmt.Println("LONGEST:", longest)
		fmt.Println("ITERATION:", i)
		fmt.Println("START:", start)
		fmt.Println("i-start+1=", i-start+1)
		fmt.Println("---------")
		//storing the variable inside of the map with a value if not present
		seen[val] = i + 1
		fmt.Println("seen[val]", seen[val])
		fmt.Println("---------")
	}
	fmt.Println("Final result", "---------")
	fmt.Println(longest)
	return longest
}

func maxVal(x, y int) int {
	if x > y {
		return x
	}
	return y
}
