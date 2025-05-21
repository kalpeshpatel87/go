package main

import "fmt"

// Factorial example using Recursive (5 = 5x4x3x2x1 = 120)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	ans := fact(5)

	fmt.Println("Factorial is :", ans)
}
