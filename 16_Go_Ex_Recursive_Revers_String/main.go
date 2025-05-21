package main

import (
	"fmt"
)

// Recursive function to reverse a string
func revString(str string) string {
	if len(str) == 0 {
		return ""
	}
	return revString(str[1:]) + string(str[0])
}

func main() {
	originalStr := "Hello Kalpesh"
	reverStr := revString(originalStr)
	fmt.Printf("The Original String is: %s\nThe Revers String is: %s\n", originalStr, reverStr)
}
