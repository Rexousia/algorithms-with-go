package main

import "fmt"

func main() {
	fmt.Println(sumRange(50))
	fmt.Println(factorial(3))
	xs := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(collectOddValues(xs))
	fmt.Println(collectOddValuesPure(xs))
}

func sumRange(num int) int {
	if num == 1 {
		return 1
	}

	return num + sumRange(num-1)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)

}

//helper method
func collectOddValues(xs []int) []int {
	//storage for odd values
	result := []int{}

	//declaring the helper function
	var helper func([]int) []int

	//defining the helper method function
	helper = func(helperInput []int) []int {

		//base case
		if len(helperInput) == 0 {
			//returning odd values when base case condition is met
			return result
		}
		//collecting odd values and storing inside of the result
		if helperInput[0]%2 != 0 {
			result = append(result, helperInput[0])
			fmt.Println("Result:", result)
		}

		//recursively deleting the first index in the slice
		return helper(append(helperInput[:0], helperInput[1:]...))
	}

	//beginning recursion
	return helper(xs)
}

//pure recursion
func collectOddValuesPure(xs []int) []int {
	newXs := []int{}
	if len(xs) == 0 {
		return newXs
	}
	if xs[0]%2 != 0 {
		newXs = append(newXs, xs[0])
	}
	finalSlice := collectOddValuesPure(append(xs[:0], xs[1:]...))
	return append(newXs, finalSlice...)
}
