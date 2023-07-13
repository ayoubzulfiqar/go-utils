package main

func pow(a int, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}
