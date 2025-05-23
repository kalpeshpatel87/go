package main

import "fmt"

type Person struct {
	Name  string
	Age   int
	Email string
}

func main() {
	fmt.Println(Person{"Kalpesh", 20, "kalpesh@gmail.com"})
	fmt.Println(Person{Name: "Ishan"})
}
