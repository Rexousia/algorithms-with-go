package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 2}
	slice2 := []int{9, 1, 4, 4}

	same(slice1, slice2)
}

func same(slice1, slice2 []int) bool {
	//checking to see if the slice length is the same
	if len(slice1) != len(slice2) {
		return false
	}
	//creating some map to count the value frequency inside of the slice
	frequencyCounter1 := make(map[int]int)
	frequencyCounter2 := make(map[int]int)

	//ranging over the first slice
	for _, val := range slice1 {
		//if the value is found inside of the map then add 1
		if _, ok := frequencyCounter1[val]; ok {
			frequencyCounter1[val] += 1
		}
		//if the value is not found inside of the map then store and set value to 1
		if _, ok := frequencyCounter1[val]; !ok {
			frequencyCounter1[val] = 1
		}
	}
	for _, val := range slice2 {
		//if the value is found inside of the map then add 1
		if _, ok := frequencyCounter2[val]; ok {
			frequencyCounter2[val] += 1

		}
		//if the value is not found inside of the map then store and set value to 1
		if _, ok := frequencyCounter2[val]; !ok {
			frequencyCounter2[val] = 1
		}
	}
	//comparing values inside of the map
	for key := range frequencyCounter1 {
		//storing exponentiation
		expoKey := key * key
		//if the key * key is not found inside of the second map then return false
		//the value is not present therefore there is no frequency pattern
		if _, ok := frequencyCounter2[expoKey]; !ok {
			return false
		}
		//if the value inside of frequency counter 1 is not equal to the value inside frequency counter 2
		if frequencyCounter2[expoKey] != frequencyCounter1[key] {
			return false
		}
	}
	fmt.Println(true)
	return true
}
