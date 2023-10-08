package main

import (
	"fmt"
	"reflect"
)

func flattenList(list interface{}) []interface{} {
	flattened := []interface{}{}

	for _, item := range list.([]interface{}) {
		if reflect.TypeOf(item).Kind() == reflect.Slice {
			flattened = append(flattened, flattenList(item)...)
		} else {
			flattened = append(flattened, item)
		}
	}

	return flattened
}

func main() {
	nestedList := []interface{}{
		[]interface{}{1, 1},
		2,
		[]interface{}{3, []interface{}{5, 8}},
	}

	flattenedList := flattenList(nestedList)

	fmt.Printf("Nested List: %#v\n", nestedList)
	fmt.Printf("Flattened List: %#v\n", flattenedList)
}
