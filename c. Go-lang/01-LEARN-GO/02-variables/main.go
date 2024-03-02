package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var age int // Declaration of the variable "Age"
	age = 23    // Initialisation of the variable "Age"
	fmt.Println("My age is", age)
	fmt.Printf("Type of variable 'age' is %T\n", age)

	// Println is used for printing a line of text without any formatting
	// Printf is used used for formatted printing.

	var name string = "ashish"
	fmt.Println("My name is ", name)
	fmt.Printf("Type of variable 'name' is %T\n", name)

	// Short declaration syntax
	car := "Wrangler"
	fmt.Println("My favourite car is ", car)
	fmt.Printf("Type of variable 'car' is %T\n", car)
	year := 2026
	fmt.Println("I'll buy this car in ", year)
	fmt.Printf("Type of variable 'year' is %T\n", year)
}
