package main

import (
	"fmt"
	"maps"
)

func main() {
	arr := make(map[string]int)
	arr["kalpesh"] = 1
	arr["karan"] = 2
	arr["kajal"] = 3

	fmt.Println("Array : ", arr)
	fmt.Println("Array of kalpesh : ", arr["kalpesh"])
	fmt.Println("Array of Length : ", len(arr))

	varOne := arr["kalpesh"]
	fmt.Println("varOne : ", varOne)

	varTwo := arr["karan"]
	fmt.Println("varTwo : ", varTwo)

	varThree := arr["kajal"]
	fmt.Println("varThree : ", varThree)

	delete(arr, "kajal")
	fmt.Println("Array after deleting kajal : ", arr)

	clear(arr)
	fmt.Println("Array after clearing : ", arr)

	arrTwo := map[string]int{
		"key1": 10,
		"key2": 20,
		"key3": 30,
	}
	fmt.Println("Array two : ", arrTwo)

	if maps.Equal(arr, arrTwo) {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}
}
