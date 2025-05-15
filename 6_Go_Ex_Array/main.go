package main

import "fmt"

func main() {
	a := [5]int{}
	fmt.Println("a:", a)

	b := [5]int{1, 2, 3, 4, 5}
	b[2] = 100
	fmt.Println("b:", b)

	xyz := [5][3]int{
		{0, 5, 10},
		{0, 15, 20},
	}
	fmt.Println("xyz:", xyz)
	fmt.Println("Length of xyz:", len(xyz))

	at := [2]string{"ABC", "DEF"}
	fmt.Println("at:", at)
}
