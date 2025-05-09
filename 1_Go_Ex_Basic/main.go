package main

import (
	"ex1/custpack"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	var str string = "This is String"
	var num int = 10
	var num2 int = 20
	var num3 int = num + num2

	fmt.Println(str)
	fmt.Println(num3)

	custpack.Printmsg()
}
