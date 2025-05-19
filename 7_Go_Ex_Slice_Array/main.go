package main

import (
	"fmt"
	"slices"
)

func main() {
	var sliceArray []string

	fmt.Println("Slice Array : ", sliceArray)
	fmt.Println("Slice Array : ", len(sliceArray))

	// Using Make

	sliceArray = make([]string, 5, 7)
	fmt.Println("Slice Array Lenght : ", len(sliceArray))
	fmt.Println("Slice Array Capacity : ", cap(sliceArray))

	sliceArray = make([]string, 2)
	fmt.Println("Slice Array Lenght : ", len(sliceArray))
	fmt.Println("Slice Array Capacity : ", cap(sliceArray))

	sliceArray = make([]string, 2, 3)
	fmt.Println("Slice Array Lenght : ", len(sliceArray))
	fmt.Println("Slice Array Capacity : ", cap(sliceArray))

	sliceArray = append(sliceArray, "A")
	sliceArray = append(sliceArray, "B")
	sliceArray = append(sliceArray, "C")
	fmt.Println("Slice Array Lenght : ", len(sliceArray))
	fmt.Println("Slice Array Capacity : ", cap(sliceArray))

	sliceArray[0] = "D"
	sliceArray[1] = "E"
	sliceArray = append(sliceArray, "F")
	sliceArray = append(sliceArray, "G")
	sliceArray = append(sliceArray, "H")
	fmt.Println("Slice Array Lenght : ", len(sliceArray))
	fmt.Println("Slice Array Capacity : ", cap(sliceArray))

	sliceArray[0] = "I"
	sliceArray[1] = "J"
	sliceArray = append(sliceArray, "K")
	sliceArray = append(sliceArray, "L")

	fmt.Println("Set Slice Array : ", sliceArray)

	// Using Copy

	copySliceArray := make([]string, 4)
	copy(copySliceArray, sliceArray)

	fmt.Println("Copy Slice Array is : ", copySliceArray)

	newArray := sliceArray[0:3]
	fmt.Println("Slice Array is : ", newArray)

	newDynamicArray := sliceArray[len(sliceArray)-3:]
	fmt.Println("Dynamic Slice Array is : ", newDynamicArray)

	arrOne := []int{1, 6, 3}
	arrTwo := []int{1, 2, 3}

	fmt.Println("All Array Equals : ", slices.Equal(arrOne, arrTwo))
}
