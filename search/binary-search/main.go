package main

import (
	"fmt"
)

func main() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 20, 22, 23, 25, 40}
	fmt.Println(binarySearch(xs, 100))
	// fmt.Println(binarySearch(xs, 6))
	// fmt.Println(binarySearch(xs, 1))
}

//divide and conquer
func binarySearch(xs []int, val int) int {
	//starting point of window
	start := 0
	//end point of window
	end := len(xs) - 1

	//assesing the values
	for start <= end {
		//
		middle := (start + end) / 2
		fmt.Println("-------------")
		fmt.Println("middel index:", middle)
		fmt.Println("middel value:", xs[middle])
		fmt.Println("-------------")
		if xs[middle] < val {
			start = middle + 1
			fmt.Println("new start value:", start)
		}
		if xs[middle] > val {
			end = middle - 1
			fmt.Println("new end value:", end)
		}
		if val == xs[middle] {
			return middle
		}
	}
	return -1
}
