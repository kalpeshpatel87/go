package main

import "fmt"

// Even Odd number with Variadic function using loop and pointer.

func checkEvenOdd(nums ...int) {
	for i := range nums {
		numPtr := &nums[i]
		if *numPtr%2 == 0 {
			fmt.Printf("Number %d is Even\n", *numPtr)
		} else {
			fmt.Printf("Number %d is Odd\n", *numPtr)
		}
	}
}

func main() {
	checkEvenOdd(1, 2, 3, 4, 5, 6)
}
