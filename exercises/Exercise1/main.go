package main

import "fmt"

func findKthElement(k int, myList []int) int {
	return myList[k]
}

func main() {
	myList := []int{1, 1, 2, 3, 5, 8}

	var k int
	fmt.Print("Enter the index of the element you want to find: ")
	fmt.Scan(&k)

	element := findKthElement(k, myList)
	fmt.Printf("The %d-th element is: %d\n", k, element)
}
