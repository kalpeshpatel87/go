package main

import (
	"fmt"
	"slices"
)

var PublicSlice []string

var privateSlice []string

func main() {

	fmt.Println("Initial Public Slice:", PublicSlice)

	PublicSlice = make([]string, 5, 7)
	fmt.Println("Public Slice Length:", len(PublicSlice))
	fmt.Println("Public Slice Capacity:", cap(PublicSlice))

	PublicSlice = make([]string, 2)
	fmt.Println("Public Slice Length:", len(PublicSlice))
	fmt.Println("Public Slice Capacity:", cap(PublicSlice))

	PublicSlice = make([]string, 2, 3)
	fmt.Println("Public Slice Length:", len(PublicSlice))
	fmt.Println("Public Slice Capacity:", cap(PublicSlice))

	PublicSlice = append(PublicSlice, "A", "B", "C")
	fmt.Println("After Appending:", PublicSlice)
	fmt.Println("Length:", len(PublicSlice))
	fmt.Println("Capacity:", cap(PublicSlice))

	PublicSlice[0] = "D"
	PublicSlice[1] = "E"
	PublicSlice = append(PublicSlice, "F", "G", "H")
	fmt.Println("After More Appending:", PublicSlice)
	fmt.Println("Length:", len(PublicSlice))
	fmt.Println("Capacity:", cap(PublicSlice))

	PublicSlice[0] = "I"
	PublicSlice[1] = "J"
	PublicSlice = append(PublicSlice, "K", "L")
	fmt.Println("Final Public Slice:", PublicSlice)

	copySlice := make([]string, 4)
	copy(copySlice, PublicSlice)
	fmt.Println("Copied Slice:", copySlice)

	subSlice := PublicSlice[0:3]
	fmt.Println("Sub Slice (0:3):", subSlice)

	lastThree := PublicSlice[len(PublicSlice)-3:]
	fmt.Println("Last 3 Elements:", lastThree)

	arrOne := []int{1, 6, 3}
	arrTwo := []int{1, 2, 3}
	fmt.Println("Are Arrays Equal:", slices.Equal(arrOne, arrTwo))

	initPrivateSlice()
	fmt.Println("Slice:", getPrivateSlice())
}

func initPrivateSlice() {
	privateSlice = []string{"Private-A", "Private-B"}
}

func getPrivateSlice() []string {
	return privateSlice
}
