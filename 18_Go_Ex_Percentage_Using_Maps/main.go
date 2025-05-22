package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	subjects := []string{"Math", "Science", "English", "Hindi", "Gujarati"}
	marks := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	totalMarks := 0

	fmt.Println("Enter marks for the following subjects (out of 100):")
	fmt.Println("--------------------------------")

	for _, subject := range subjects {
		fmt.Printf("Enter Marks For %s: ", subject)
		scanner.Scan()
		input := scanner.Text()
		mark, err := strconv.Atoi(input)
		if err != nil || mark < 0 || mark > 100 {
			fmt.Println("Invalid input. Please enter a number between 0 and 100.")
			return
		}
		marks[subject] = mark
		totalMarks += mark
	}

	percentage := float64(totalMarks) / float64(len(subjects))
	fmt.Println("--------------------------------")
	fmt.Println("Result Summary:")
	fmt.Println("--------------------------------")
	for subject, mark := range marks {
		fmt.Printf("%s: %d\n", subject, mark)
	}
	fmt.Println("--------------------------------")
	fmt.Printf("Total Obtained Marks: %d / %d\n", totalMarks, len(subjects)*100)
	fmt.Println("--------------------------------")
	fmt.Printf("Your Percentage: %.2f%%\n", percentage)

	var finalResult string
	if percentage >= 70 {
		finalResult = "Congratulations You Are Passed With First Class"
	} else if percentage >= 50 {
		finalResult = "Congratulations You Are Passed With Second Class"
	} else if percentage >= 35 {
		finalResult = "Congratulations You Are Passed With Third Class"
	} else {
		finalResult = "Sorry You Are FAIL Batter Luck Next Time"
	}

	fmt.Println("--------------------------------")
	fmt.Printf("Final Result: %s\n", finalResult)
	fmt.Println("--------------------------------")
}
