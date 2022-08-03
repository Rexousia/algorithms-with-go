package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binarySearch(xs, 7))
	fmt.Println(binarySearch(xs, 6))
	fmt.Println(binarySearch(xs, 1))
}

func binarySearch(xs []int, val int) int {
	start := 0
	end := len(xs) - 1

	for start <= end {
		middle := (start + end) / 2
		if xs[middle] < val {
			start = middle + 1
		}
		if xs[middle] > val {
			end = middle - 1
		}
		if val == xs[middle] {
			return middle
		}
	}
	return -1
}
