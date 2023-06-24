package main

/*

The `copyFrom` function creates a new map and copies all the
key-value pairs from the original map to the new map, effectively creating a 
duplicate map with the same key-value associations.


*/

func copyFrom(array map[int]int) map[int]int {
	newArray := make(map[int]int)
	for key, value := range array {
		newArray[key] = value
	}

	return newArray
}
