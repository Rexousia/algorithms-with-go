package main

import (
	"fmt"
)

func main() {
	xs := []int{2, 6, 9, 2, 1, 8, 5, 6, 3}
	fmt.Println(xs)
	maxSubsliceSum(xs, 3)
}

func maxSubsliceSum(xs []int, num int) int {
	maxSum := 0
	tempSum := 0
	if len(xs) < num {
		return 0
	}

	for i := 0; i < num; i++ {
		maxSum += xs[i]
		fmt.Printf("Position: \t %d \nxs[i]:\t\t %d\n", i, xs[i])
		fmt.Printf("Position: \t %d \nmaxSum inside first loop:\t %d\n", i, maxSum)
	}

	tempSum = maxSum
	fmt.Printf("\ntempSum = maxSum outside of loop: \t %d\n", tempSum)

	for i := num; i < len(xs); i++ {
		fmt.Println("tempSum - xs[i-num] + xs[i]:")
		fmt.Printf("%d - %d + %d\n", tempSum, xs[i-num], xs[i])
		fmt.Printf("Second loop tempSum:  \t %d \n", tempSum)
		tempSum = tempSum - xs[i-num] + xs[i]
		fmt.Printf("Second loop xs[i-num]: \t %d \n", xs[i-num])
		fmt.Printf("Second loop xs[i]:\t %d \n", xs[i])
		maxSum = Max(maxSum, tempSum)
		fmt.Printf("Comparing tempSum and newMaxSum: %d \n\n", maxSum)
	}
	return maxSum
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
