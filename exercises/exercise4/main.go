package main

import "fmt"

func eliminateDupes(inputList []int) []int {
	if len(inputList) <= 1 {
		return inputList
	}

	result := []int{inputList[0]}

	for i := 1; i < len(inputList); i++ {
		if inputList[i] != inputList[i-1] {
			result = append(result, inputList[i])
		}
	}

	return result
}

func main() {
	myList := []int{1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 6, 6}
	result := eliminateDupes(myList)

	fmt.Printf("%v\n", myList)
	fmt.Printf("Consecutive Duplicates Eliminated: %v\n", result)
}
