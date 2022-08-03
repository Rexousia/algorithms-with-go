package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "cat"
	fmt.Println(reverse(str))
}

func reverse(str string) string {
	if len(str) <= 0 {
		return str
	}

	newStr := strings.Split(str, "")
	rvr := []string{}
	var helper func(str1 []string) string
	helper = func(str1 []string) string {
		if len(str1) == 0 {
			//joining the entire slice
			return strings.Join(rvr[:], "")
		}
		//adding the last index or value of the slice to rvr
		rvr = append(rvr, str1[len(str1)-1:]...)
		//removing the last letter of slice
		return helper(str1[:len(str1)-1])
	}
	return helper(newStr)
}
