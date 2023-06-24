package main

/*

The provided code defines a function called `copyFrom` that creates a new copy of a given `map[int]int` (a map with integer keys and integer values).

Here's an explanation of how the code works:

1. The function `copyFrom` takes a single argument, `array`, which is the map that needs to be copied.

2. Inside the function, a new map, `newArray`, is created using the `make` function. This initializes an empty map.

3. The function then iterates over each key-value pair in the original `array` using a `for` loop and the `range` keyword.

4. In each iteration, the key and value of the current key-value pair are assigned to the variables `key` and `value`, respectively.

5. Inside the loop, the line `newArray[key] = value` copies the key-value pair from the original `array` to the `newArray` map. This line creates a new entry in `newArray` with the same key and value as the corresponding entry in the original `array`.

6. After iterating over all the key-value pairs in `array`, the `newArray` map is fully populated with the copied entries.

7. Finally, the `newArray` map is returned as the result of the `copyFrom` function.

In summary, the `copyFrom` function creates a new map and copies all the key-value pairs from the original map to the new map, effectively creating a duplicate map with the same key-value associations.

*/

func copyFrom(array map[int]int) map[int]int {
	newArray := make(map[int]int)
	for key, value := range array {
		newArray[key] = value
	}

	return newArray
}
