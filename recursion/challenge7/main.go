package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("tacocat"))
	fmt.Println(isPalindrome("awesome"))

}
func isPalindrome(str string) bool {
	newStr := strings.Split(str, "")
	rvr := []string{}
	var helper func(str1 []string) bool

	helper = func(str1 []string) bool {

		//checking and comparing the length of rvr and newStr
		if len(rvr) == len(newStr) {
			//checking to see if palindrome
			if strings.Join(rvr[:], "") == str {
				return true
			}
			//checking to see if not palindrome
			if strings.Join(rvr[:], "") != str {
				return false
			}
		}

		//adding the last letter from str1 to rvr
		rvr = append(rvr, str1[len(str1)-1:]...)
		fmt.Println(rvr)

		//popping off the last letter of str1 and running recursively
		return helper(str1[:len(str1)-1])
	}
	return helper(newStr)
}
