package main


/*

The `abs` function takes an integer `num` as input and returns its absolute value.
If the input number is negative, the function returns the negation of that number.
Otherwise, if the input number is non-negative, the function returns the number itself.
In other words, the `abs` function ensures that the returned value is the positive
magnitude of the input number, disregarding its sign.

*/

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}