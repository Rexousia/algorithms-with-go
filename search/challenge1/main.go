package main

func main() {
	arr := [3]int{2, 3, 1}
	Gimme(arr)
}

func Gimme(array [3]int) int {
	if (array[0] > array[1]) != (array[0] > array[2]) {
		return 0
	} else if (array[1] > array[0]) != (array[1] > array[2]) {
		return 1
	} else {
		return 2
	}
}
