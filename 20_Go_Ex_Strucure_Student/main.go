package main

import (
	"fmt"
)

type StudentSubject struct {
	Math    int
	Science int
	English int
}

type Student struct {
	Name     string
	Age      int
	RollNo   int
	Subjects StudentSubject
}

func (s Student) TotalMarks() int {
	return s.Subjects.Math + s.Subjects.Science + s.Subjects.English
}

func (s Student) AverageMarks() float64 {
	return float64(s.TotalMarks()) / 3
}

func main() {
	student := Student{
		Name:   "Kalpesh",
		Age:    30,
		RollNo: 87,
		Subjects: StudentSubject{
			Math:    85,
			Science: 90,
			English: 78,
		},
	}
	fmt.Println("--------------------------------")
	fmt.Printf("Student Information:\n")
	fmt.Println("--------------------------------")
	fmt.Printf("Student Name: %s\n", student.Name)
	fmt.Printf("Age: %d\n", student.Age)
	fmt.Printf("Roll No: %d\n", student.RollNo)
	fmt.Println("--------------------------------")
	fmt.Println("Student Result Summary:")
	fmt.Println("--------------------------------")
	fmt.Printf("Math: %d\nScience: %d\nEnglish: %d\n", student.Subjects.Math, student.Subjects.Science, student.Subjects.English)
	fmt.Println("--------------------------------")
	fmt.Printf("Total Marks: %d\n", student.TotalMarks())
	fmt.Println("--------------------------------")
	fmt.Printf("Average Marks: %.2f\n", student.AverageMarks())
	fmt.Println("--------------------------------")
}
