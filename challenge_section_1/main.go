package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sameFrequency(789610, 109787)
}

func sameFrequency(arg1, arg2 int) bool {
	st1 := strconv.Itoa(arg1)
	st2 := strconv.Itoa(arg2)

	if len(st1) != len(st2) {
		fmt.Println(false)
		return false
	}

	newS1 := strings.Split(st1, "")
	newS2 := strings.Split(st2, "")

	frequencyMap1 := map[string]int{}
	frequencyMap2 := map[string]int{}

	for _, val := range newS1 {
		if _, ok := frequencyMap1[val]; ok {
			frequencyMap1[val] += 1
		}

		if _, ok := frequencyMap1[val]; !ok {
			frequencyMap1[val] = 1
		}

	}
	for _, val := range newS2 {
		if _, ok := frequencyMap2[val]; ok {
			frequencyMap2[val] += 1
		}

		if _, ok := frequencyMap2[val]; !ok {
			frequencyMap2[val] = 1
		}

	}

	for key := range frequencyMap1 {
		if frequencyMap1[key] != frequencyMap2[key] {
			fmt.Println(false)
			return false
		}

		if _, ok := frequencyMap2[key]; !ok {
			fmt.Println(false)
			return false
		}
	}
	fmt.Println(true)
	return true
}
