package main

////////////////////////////////////////////////
//Time Complexity o(n)/////////////////////////
//////////////////////////////////////////////

import (
	"fmt"
	"strings"
)

func main() {
	sameString("iceman", "cinem4")
}

func sameString(string1, string2 string) bool {
	//Checking for the length of the string
	if len(string1) != len(string2) {
		return false
	}
	//Creating some maps to store the values of the string inside of and count frequency
	frequencyString1 := map[string]int{}
	frequencyString2 := map[string]int{}

	//Splitting the string
	s1 := strings.Split(string1, "")
	s2 := strings.Split(string2, "")

	//ranging over the string
	for _, val := range s1 {
		//if the value is present inside of the map add one
		if _, ok := frequencyString1[val]; ok {
			frequencyString1[val] += 1
		}
		//if the value is not present store the value in the map and set the value to one
		if _, ok := frequencyString1[val]; !ok {
			frequencyString1[val] = 1
		}
	}

	for _, val := range s2 {
		//if the value is present inside of the map add one
		if _, ok := frequencyString2[val]; ok {
			frequencyString2[val] += 1
		}
		//if the value is not present store the value in the map and set the value to one
		if _, ok := frequencyString2[val]; !ok {
			frequencyString2[val] = 1
		}
	}

	for key := range frequencyString1 {
		//if the value in frequencystring1 is not present inside of frequencystring2 return false
		if _, ok := frequencyString2[key]; !ok {
			fmt.Println(false)
			return false
		}
		//if the value inside of the frequencystring1 is not equal to frequencystring2 return false
		if frequencyString1[key] != frequencyString2[key] {
			fmt.Println(false)
			return false
		}
	}
	fmt.Println("s1", frequencyString1)
	fmt.Println("s2", frequencyString2)
	fmt.Println(true)
	return true
}
