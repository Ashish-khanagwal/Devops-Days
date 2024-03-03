package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's learn about taking user input")

	// This is how we take input from the user
	var firstName string
	// fmt.Print("Enter your name: ")

	// We used "&" --> Because this is a pointer which points to the memory address of the declared variable
	// hence, this is known as passing by reference.
	// fmt.Scanln(&firstName)
	// fmt.Println("Hello", firstName)

	// Multiple inputs
	var lastName string
	var age int
	fmt.Print("Enter your firstname, lastname, and age: ")
	fmt.Scanln(&firstName, &lastName, &age)
	fmt.Printf("Your name is %v %v and you're %v years old\n", firstName, lastName, age)
}
