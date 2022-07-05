package main

import "fmt"

func main() {
	d2 := []int{1, 4, 2, 10, 23, 3, 1, 0, 20}
	d1 := []int{100, 200, 300, 400}
	d3 := []int{-3, 4, 0, -2, 6, -1}
	d4 := []int{3, -2, 7, -4, 1, -1, 4, -2, 1}
	d5 := []int{2, 3}
	maxSubarraySum(d1, 2)
	maxSubarraySum(d2, 4)
	maxSubarraySum(d3, 2)
	maxSubarraySum(d4, 2)
	maxSubarraySum(d5, 3)
}

func maxSubarraySum(xs []int, num int) int {
	tempSum := 0
	maxSum := 0
	if len(xs) < num {
		return 0
	}
	for i := 0; i < num; i++ {
		maxSum += xs[i]
	}
	tempSum = maxSum
	for i := num; i < len(xs); i++ {
		tempSum = tempSum - xs[i-num] + xs[i]
		maxSum = max(tempSum, maxSum)

		fmt.Println(maxSum)
	}
	return maxSum
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
