package main

import (
	"fmt"
)

type Role int

const (
	Admin Role = iota
	Local
	User
)

func (r Role) String() string {
	return [...]string{"Admin", "Local", "User"}[r]
}

func main() {
	var myRole Role = Local

	fmt.Println("Your role is:", myRole)

	switch myRole {
	case Admin:
		fmt.Println("You have full access.")
	case Local:
		fmt.Println("You have local access.")
	case User:
		fmt.Println("You have limited access.")
	}
}
