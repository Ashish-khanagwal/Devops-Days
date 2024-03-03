package main

import "fmt"

func main() {
	fmt.Println("Here, we will discuss some advanced things in golang")

	var name = "Ashish"

	// We can use "name" variable in multiple ways.
	fmt.Println("My name is", name, "I'm an engineer")  // first way
	fmt.Printf("My name is %v I'm an Engineer\n", name) // second way

	// Printf --> tells go to format the variabels we are using

	age := 23
	firstName := "Ashish"
	lastName := "Khanagwal"

	// %s is a placeholder for string and %d is placeholder for integer
	fmt.Printf("My name is %s and I'm %d years old\n", firstName, age)
	fmt.Printf("My name is %s %s and I'm %d years old\n", firstName, lastName, age)
}
