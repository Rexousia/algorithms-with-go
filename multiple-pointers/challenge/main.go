package main

import "fmt"

func main() {
	xs := []int{1, 1, 1, 1, 1, 1, 2}
	xs2 := []int{1, 2, 3, 4, 4, 4, 7, 7, 12, 12, 13}
	xs3 := []int{}
	xs4 := []int{-2, -1, -1, 0, 1}
	countUniqueValues(xs)
	countUniqueValues(xs2)
	countUniqueValues(xs3)
	countUniqueValues(xs4)
	countUniqueValuesWithPointers(xs)
	countUniqueValuesWithPointers(xs2)
	countUniqueValuesWithPointers(xs3)
	countUniqueValuesWithPointers(xs4)

}

//with pointers o(n)
func countUniqueValuesWithPointers(xs []int) int {
	if len(xs) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(xs); j++ {
		if xs[i] != xs[j] {
			i++
			xs[i] = xs[j]
		}
	}
	fmt.Println(i + 1)
	return i + 1
}

//without pointers o(n)
func countUniqueValues(xs []int) int {
	uniqueCharacters := map[int]int{}
	for _, val := range xs {
		if _, ok := uniqueCharacters[val]; !ok {
			uniqueCharacters[val] = 1
		}
	}
	fmt.Println("Number of unique characters", len(uniqueCharacters))
	return len(uniqueCharacters)
}
