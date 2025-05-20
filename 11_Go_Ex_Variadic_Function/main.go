package main

import "fmt"

func sum(num ...int) {
	fmt.Print(num, " ")

	total := 0

	for _, val := range num {
		total = total + val
	}

	fmt.Println("Total : ", total)
}
func main() {
	sum(2, 4, 6)
}
