package main

import (
	"fmt"
)

func main() {
	s1 := "wowomgomlomg"
	s2 := "omg"
	fmt.Println(naiveSearch(s1, s2))
	s3 := "lorie loled"
	s4 := "lo"
	fmt.Println(naiveSearch2(s3, s4))
}

//first way
func naiveSearch(s1, s2 string) int {
	tempCounter := 0
	mainCounter := 0
	for i := 0; i < len(s1); i++ {

		for j := 0; j < len(s2); j++ {
			fmt.Println(s1[i], s2[j])
			if s1[i] != s2[j] {
				tempCounter = 0
				break
			}
			if s1[i] == s2[j] {
				tempCounter++
				i++
			}
			if tempCounter == len(s2) {
				mainCounter++
				tempCounter = 0
			}
		}

	}
	return mainCounter
}

//second way
func naiveSearch2(long, short string) int {
	mainCounter := 0
	for i := 0; i < len(long); i++ {

		for j := 0; j < len(short); j++ {
			if long[i+j] != short[j] {
				break
			}
			if j == len(short)-1 {
				mainCounter++
			}
		}
	}
	return mainCounter
}
