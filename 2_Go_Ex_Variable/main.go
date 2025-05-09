package main

import "fmt"

func printMainMsg() {
	fmt.Println("This is calling from the Main package.")
}

func main() {
	var PublicVar string = "This is Public Variable"   //This is Public Variable.
	var privateVar string = "This is Private Variable" //This is Private Variable.

	fmt.Println("PublicVar : ", PublicVar)
	fmt.Println("privateVar : ", privateVar)

	printMainMsg()

	// Short Declaration Variable

	sVar1, sVar2, sVar3 := "Variable 1", "Variable 2", "Variable 1"

	fmt.Println("Short Declaration Variable Example : ", sVar1, sVar2, sVar3)

	// Prin Formator
	// For String = %s
	// For Int = %d
	// For Float = %f

	name := "Kalpesh"
	salary := 26500.956
	years := 5

	fmt.Printf("%s has %d years of experience and his salary is %.2f \n", name, years, salary)

}
