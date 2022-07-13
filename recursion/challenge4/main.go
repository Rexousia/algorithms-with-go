package main

import "fmt"

func main() {
	fmt.Println(recursiveRange(6))
	fmt.Println(recursiveRange(10))
}

//recursiveRange
func recursiveRange(num int) int {
	if num == 1 {
		return 1
	}
	return num + recursiveRange(num-1)
}
