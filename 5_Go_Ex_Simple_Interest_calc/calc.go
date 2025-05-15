package main

import (
	"fmt"

	"example.com/simpleinterest/calculation"
)

func getInput(text string) (inputVal float64) {
	fmt.Println(text)
	fmt.Scan(&inputVal)
	return
}

func main() {
	var principal float64 = getInput("Enter Principal: ")
	var roi float64 = getInput("Enter ROI: ")
	var tenor float64 = getInput("Enter Tenor: ")

	var result = calculation.CalcInterest(principal, roi, tenor)

	fmt.Printf("Simple Interest is : %.2f\n", result)
}
