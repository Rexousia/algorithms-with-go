package main

import "fmt"

func main() {
	xs := []int{10, 4, 1, 2, 4, 7, 9}
	fmt.Println(selectionSort(xs))
}

func selectionSort(xs []int) []int {

	for i := 0; i < len(xs); i++ {
		minVal := i
		for j := i + 1; j < len(xs); j++ {
			fmt.Println(xs, minVal, xs[i], xs[j])
			if xs[j] < xs[minVal] {
				minVal = j
			}

		}
		if i != minVal {
			xs[i], xs[minVal] = xs[minVal], xs[i]
		}

	}
	return xs
}
