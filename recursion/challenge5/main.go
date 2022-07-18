package main

import "fmt"

func main() {

	fmt.Println(fib(4))
	fmt.Println(fib(10))
	fmt.Println(fib(28))
	fmt.Println(fib(35))
}

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
