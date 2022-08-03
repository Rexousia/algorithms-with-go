package main

import (
	"fmt"
	"strings"
)

func main() {
	xs := []string{"car", "taco", "banana"}
	fmt.Println(capitalizeFirst(xs))
}

func capitalizeFirst(xs []string) []string {
	newString := []string{}

	var helper func(xs []string) []string
	helper = func(xs []string) []string {
		if len(xs) == 0 {
			return newString
		}

		chara := strings.ToUpper((string(xs[0][0])))
		newString = append(newString, chara+xs[0][1:])
		return helper(xs[1:])

	}
	return helper(xs)

}
