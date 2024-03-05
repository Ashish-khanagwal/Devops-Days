package main

import "fmt"

// main is the entry point of the program
func main() {
	fmt.Println("Now we will learn about functions")
	greet() // calling the function (1)

	greetWithName("Ashish") // (2)

	fmt.Println(add(2, 3)) // (3)
}

// greet is a function that prints "Hello, World!" to the console (1)
func greet() {
	fmt.Println("Hello, World!")
}

// greetWithName is a function that prints "Hello, name" to the console (2)
func greetWithName(name string) {
	fmt.Println("Hello, ", name)
}

// Function with return type (3)
func add(a, b int) int {
	return a + b
}
