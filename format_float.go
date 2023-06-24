package main

/*
 Format float without any package - no stringConversion No fmt package


*/

func formatFloat(value float64) string {
	intPart := int(value)
	fracPart := int((value - float64(intPart)) * 100)
	return itoa(intPart) + "." + itoa(fracPart)
}

func itoa(n int) string {
	digits := "0123456789"
	if n == 0 {
		return "0"
	}
	var result string
	for n > 0 {
		digit := n % 10
		result = string(digits[digit]) + result
		n /= 10
	}
	return result
}
