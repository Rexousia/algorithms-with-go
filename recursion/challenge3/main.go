package main

import "fmt"

func main() {

	fmt.Println(productOfArray([]int{1, 2, 3}))
	fmt.Println(productOfArray([]int{1, 2, 3, 10}))
}

//returning int
func productOfArray(xs []int) int {
	total := 1

	//returning int
	var helper func(xs2 []int) int

	//returning int
	helper = func(xs2 []int) int {
		//returning the total
		if len(xs2) == 0 {
			return total
		}

		//adding the value at index 0 to the total
		if xs2[0] > 0 {
			total *= xs2[0]
		}

		//deleting the first value in the slice
		return helper(append(xs2[:0], xs2[1:]...))
	}

	//return the int
	return helper(xs)
}
