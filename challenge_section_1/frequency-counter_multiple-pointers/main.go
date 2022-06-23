package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(areThereDuplicateInt(1, 2, 3, 4, 5, 5, 7))
	fmt.Println(areThereDuplicateInt(1, 2, 4, 5, 6))
	fmt.Println(areThereDuplicateString("a", "b", "c", "d", "a"))
	fmt.Println(areThereDuplicateMultiplePointers(1, 2, 3, 4, 5, 6))
	fmt.Println(areThereDuplicateMultiplePointers(1, 2, 3, 4, 4, 5))
	fmt.Println(areThereDuplicateMultiplePointers(1, 2, 2))
	fmt.Println(areThereDuplicateMultiplePointers(1, 2, 7, 8, 3, 5, 4, 6, 9, 1))

}

func areThereDuplicateInt(arg ...int) bool {
	frequencyCounter := map[int]int{}

	for _, val := range arg {
		if _, ok := frequencyCounter[val]; ok {
			frequencyCounter[val] += 1
		}
		if _, ok := frequencyCounter[val]; !ok {
			frequencyCounter[val] = 1
		}
	}

	for _, i := range frequencyCounter {
		if i >= 2 {
			return true
		}
	}
	return false
}

func areThereDuplicateString(arg ...string) bool {
	frequencyCounter := map[string]int{}

	for _, val := range arg {
		if _, ok := frequencyCounter[val]; ok {
			frequencyCounter[val] += 1
		}
		if _, ok := frequencyCounter[val]; !ok {
			frequencyCounter[val] = 1
		}
	}

	for _, i := range frequencyCounter {
		if i >= 2 {
			return true
		}
	}
	return false
}

func areThereDuplicateMultiplePointers(arg ...int) bool {
	sort.Ints(arg)
	start := 0
	next := 1

	for next < len(arg) {
		if arg[start] == arg[next] {
			return true
		}
		start++
		next++
	}
	return false
}
