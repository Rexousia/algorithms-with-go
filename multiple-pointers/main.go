package main

import "fmt"

func main() {
	random := []int{-4, -3, -2, -1, 0, 1, 2, 5}
	random2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sumZero(random)
	sumZero(random2)
}

func sumZero(slice3 []int) []int {
	newSlice := []int{}
	var left = 0
	var right = len(slice3) - 1
	for left < right {
		var sum = slice3[left] + slice3[right]
		if sum == 0 {
			newSlice = []int{slice3[left], slice3[right]}
			fmt.Println("This is inside of the for loop", newSlice)
			return newSlice
		}
		if sum > 0 {
			right--
		}
		if sum < 0 {
			left++
		}
	}
	fmt.Println("This is outside of the for loop", newSlice)
	return newSlice
}
