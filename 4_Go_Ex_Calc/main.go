package main

import "fmt"

func main() {
	var Operand1, Operand2 float64

	fmt.Print("Enter First Number: ")
	fmt.Scan(&Operand1)

	fmt.Print("Enter Second Number: ")
	fmt.Scan(&Operand2)

	addition := Operand1 + Operand2
	substraction := Operand1 - Operand2
	multiplication := Operand1 * Operand2
	division := Operand1 / Operand2

	fmt.Printf("The Addition is: %.2f\n", addition)
	fmt.Printf("The Substraction is: %.2f\n", substraction)
	fmt.Printf("The Multiplication is: %.2f\n", multiplication)
	fmt.Printf("The Division is: %.2f\n", division)
}
