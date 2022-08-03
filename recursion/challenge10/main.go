package main

import (
	"fmt"
	"strings"
)

func main() {
	xs := []string{"i", "am", "learning", "recursion"}

	fmt.Println(capitalizeWord(xs))
}

func capitalizeWord(xs []string) []string {
	newString := []string{}

	var helper func(xs []string) []string
	helper = func(xs []string) []string {
		if len(xs) == 0 {
			return newString
		}
		newString = append(newString, strings.ToUpper(xs[0]))
		return helper(xs[1:])
	}
	return helper(xs)
}
