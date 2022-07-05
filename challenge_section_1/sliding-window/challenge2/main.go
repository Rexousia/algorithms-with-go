package main

import "fmt"

func main() {
	d2 := []int{2, 3, 1, 2, 4, 3}
	d1 := []int{2, 1, 6, 5, 4}
	d3 := []int{3, 1, 7, 11, 2, 9, 8, 21, 62, 33, 19}
	d4 := []int{1, 4, 16, 22, 5, 7, 8, 9, 10}
	d5 := []int{1, 4, 16, 22, 5, 7, 8, 9, 10}
	minSubArrayLen(d1, 7)
	minSubArrayLen(d2, 9)
	minSubArrayLen(d3, 52)
	minSubArrayLen(d4, 39)
	minSubArrayLen(d5, 55)
}

func minSubArrayLen(nums []int, s int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right, res, sum := 0, -1, n+1, 0
	for left < n {
		if (right+1) < n && sum < s {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= s {
			res = min(res, right-left+1)
		}
	}
	if res == n+1 {
		return 0
	}
	fmt.Println(res)
	return res
}
func min(a int, b int) int {
	if a > b {
		return b
	}

	return a
}
