package main

import "fmt"

func main() {
	isSubsequence("hello", "hello world")
	isSubsequence("sing", "sting")
	isSubsequence("abc", "abracadabra")
	isSubsequence("abc", "acb")
}

func isSubsequence(s1, s2 string) bool {

	//xs1 = 5
	//xs2 = 11
	var xs1, xs2 = len(s1), len(s2)
	fmt.Println("--------------------")
	fmt.Println(xs1)
	fmt.Println(xs2)
	fmt.Println("--------------------")
	if xs1 > xs2 {
		fmt.Printf("false")
		return false
	}
	fmt.Println(s1)
	fmt.Println(s2)
	var i, j = 0, 0
	for i < xs1 && j < xs2 {
		fmt.Println("s1[i]", s1[i])
		fmt.Println("s2[j]", s2[j])
		if s1[i] == s2[j] {
			i += 1
		}
		j += 1
	}
	fmt.Println(i == xs1)
	return i == xs1
}
