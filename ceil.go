package main

func ceil(x, y int) int {
	div, rem := x/y, x%y
	if rem > 0 {
		div++
	}
	return div
}