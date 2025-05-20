package main

import "fmt"

func main() {
	var i = 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}
	// Using Range
	for k := range 5 {
		fmt.Println("Range: ", k)
	}
	// Using Break
	for {
		fmt.Println("Continue Loop...")
		break
	}
	// Print Odd Number but even number continue in loop
	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
