package main

import (
	"fmt"
)

// Recursive function to sum the digits of a number
func sumDigits(n int) int {
	if n == 0 {
		return 0
	}
	return n%10 + sumDigits(n/10)
}

func main() {
	number := 1505
	result := sumDigits(number)
	fmt.Printf("The sum of digits of %d is %d\n", number, result)
}
