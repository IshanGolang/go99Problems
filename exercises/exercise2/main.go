package main

import "fmt"

func reversedList(list []int) []int {
	length := len(list)
	reversed := make([]int, length)

	for i, val := range list {
		reversed[length-i-1] = val
	}

	return reversed
}

func main() {
	myList := []int{1, 1, 2, 3, 5, 8}
	reversedList := reversedList(myList)

	fmt.Printf("%v\n", myList)
	fmt.Printf("Reversed List: %v\n", reversedList)
}
