package main

import (
	"fmt"
)

func isEqual[T any, U any](list []T, other []U) bool {
	if len(list) != len(other) {
		return false
	}

	for i := 0; i < len(list); i++ {
		if fmt.Sprintf("%v", list[i]) != fmt.Sprintf("%v", other[i]) {
			return false
		}
	}

	return true
}

func main() {
	stringList := []string{"one", "two", "three"}
	intList := []int{1, 2, 3}

	equal := isEqual(stringList, intList)
	fmt.Println(equal) // Output: true
}
