package main

import "fmt"

const (
	// Enumeration of constants
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

func main() {
	fmt.Println("Learn about constants")

	// Simple explanation

	const pi = 3.14
	fmt.Println("Value of pi is", pi)

	// Typed and Untyped constants

	const typed int = 56 // Explicitly defines the type of number
	const untyped = 56
	fmt.Println("Value of typed constant is", typed)
	fmt.Println("Value of untyped constant is", untyped)
	fmt.Printf("Type of untyped constant is %T\n", untyped)

	// Enumeration of constants
	// enumeration in programming means giving names to related things so we can work with them easily and remember them better in our code.

	fmt.Println("Days of the week")
	fmt.Println("Sunday", Sunday)
	fmt.Println("Monday", Monday)
	fmt.Println("Tuesday", Tuesday)
	fmt.Println("Wednesday", Wednesday)
	fmt.Println("Thursday", Thursday)
	fmt.Println("Friday", Friday)
	fmt.Println("Saturday", Saturday)
}
