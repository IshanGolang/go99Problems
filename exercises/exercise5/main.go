package main

import (
	"fmt"
)

func packConsecutiveDuplicates(inputList []interface{}) [][]interface{} {
	if len(inputList) == 0 {
		return nil
	}

	packedList := [][]interface{}{}
	currentGroup := []interface{}{inputList[0]}

	for i := 1; i < len(inputList); i++ {
		if inputList[i] == inputList[i-1] {
			currentGroup = append(currentGroup, inputList[i])
		} else {
			packedList = append(packedList, currentGroup)
			currentGroup = []interface{}{inputList[i]}
		}
	}

	packedList = append(packedList, currentGroup) // Add the last group

	return packedList
}

func main() {
	myList := []interface{}{'a', 'a', 'a', 'a', 'b', 'c', 'c', 'a', 'a', 'd', 'e', 'e', 'e', 'e'}
	packedList := packConsecutiveDuplicates(myList)

	fmt.Printf("Original List: %v\n", myList)
	fmt.Printf("Packed List: %v\n", packedList)
}
