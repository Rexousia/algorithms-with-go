package main

import "fmt"

func main() {
	xs := []int{2, 7, 6, 8, 5, 10}
	fmt.Println(bubbleSort(xs))
}

//optomized with no swaps
func bubbleSort(array []int) []int {
	var noSwaps bool
	for i := 0; i < len(array)-1; i++ {
		noSwaps = true
		for j := 0; j < len(array)-i-1; j++ {
			fmt.Println("inner loop condition:", len(array)-i-1)
			if array[j] > array[j+1] {
				//swapping values
				array[j], array[j+1] = array[j+1], array[j]
				noSwaps = false
			}
			if noSwaps {
				break
			}
		}
	}
	return array
}
