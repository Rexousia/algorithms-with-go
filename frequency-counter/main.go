package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 2}
	arr2 := []int{9, 1, 4, 4}

	same(arr1, arr2)
}

func same(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	frequencyCounter1 := make(map[int]int)
	frequencyCounter2 := make(map[int]int)
	for _, val := range arr1 {
		if _, ok := frequencyCounter1[val]; ok {
			frequencyCounter1[val] += 1
		}
		if _, ok := frequencyCounter1[val]; !ok {
			frequencyCounter1[val] = 1
		}
	}
	for _, val := range arr2 {
		if _, ok := frequencyCounter2[val]; ok {
			frequencyCounter2[val] += 1

		}
		if _, ok := frequencyCounter2[val]; !ok {
			frequencyCounter2[val] = 1
		}
	}
	for key := range frequencyCounter1 {
		expoKey := key * key
		if _, ok := frequencyCounter2[expoKey]; !ok {
			return false
		}
		if frequencyCounter2[expoKey] != frequencyCounter1[key] {
			return false
		}
	}
	fmt.Println(true)
	return true
}
