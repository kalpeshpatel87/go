package main

import "fmt"

func newValue(ival int) {
	ival = 0
}

func newValuePointer(ival *int) {
	*ival = 0
}
func main() {
	i := 1
	fmt.Println("I : ", i)

	newValue(i)
	fmt.Println("I : ", i)

	newValuePointer(&i)
	fmt.Println("I : ", i)
}
