package main

import "fmt"

func main() {
	data1 := []int{1, 2, 3}
	data2 := []int{1, 3, 3, 5, 6, 7, 10, 12, 19}
	data3 := []int{-1, 0, 3, 4, 5, 6}
	data4 := []int{}

	fmt.Println("-------------")
	fmt.Println(avearagePair(data1, 2.5))
	fmt.Println("-------------")
	fmt.Println(avearagePair(data2, 8))
	fmt.Println("-------------")
	fmt.Println(avearagePair(data3, 4.1))
	fmt.Println("-------------")
	fmt.Println(avearagePair(data4, 4))
	fmt.Println("-------------")
}

func avearagePair(arg []int, avg float64) bool {
	if len(arg) == 0 {
		return false
	}
	left := 0
	right := 1

	for left < right {
		sum := arg[left] + arg[right]
		// fmt.Println("Arg[left]:", arg[left])
		// fmt.Println("Arg[right]:", arg[right])
		// fmt.Println("Sum:", sum)
		newAvg := float64(sum) / 2
		// fmt.Println("newAvg:", newAvg)

		if newAvg == avg {
			return true
		}

		if right == len(arg)-1 && newAvg != avg {
			left++
			if left == len(arg)-1 {
				break
			}
			right = left + 1
		}
		if right != len(arg)-1 {
			right++
		}

	}
	return false
}
