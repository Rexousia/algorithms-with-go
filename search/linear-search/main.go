package main

import "fmt"

func main() {
	xs := []int{1, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println(linearSearch(10, xs))
	fmt.Println(linearSearch(12, xs))

}

func linearSearch(val int, xs []int) int {
	for i, num := range xs {
		if num == val {
			return i
		}
	}
	return -1
}
