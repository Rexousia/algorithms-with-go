package main

import "fmt"

func main() {
	arr := []int{2, 1, 9, 76, 4}
	fmt.Println(insertionSort(arr))

}

func insertionSort(a []int) []int {
	// creating a new slice of a size equal to the original
	result := make([]int, len(a))
	// copying elements from the original slice to the new slice
	copy(result, a)
	// iterating through all elements of the slice starting with the second
	for i := 1; i < len(result); i++ {
		// key holds the current value
		key := result[i]
		// j holds the index of the previous value
		j := i - 1
		// we iterate backwards starting with the previous value and compare it with the current value
		for j >= 0 && result[j] > key {
			// if the previous value is greater then we move it forward one position
			result[j+1] = result[j]
			// we decrement the index for the previous value so that we check the next value backwards
			j--
		}
		// at the end of the inner loop we want to potentially move the current value of the main loop to a new position (if anything was shifted above)
		result[j+1] = key
	}
	// returning the sorted slice
	return result
}
