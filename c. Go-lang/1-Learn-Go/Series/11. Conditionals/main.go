package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's learn about conditionals")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("what is our topic today?: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("We will be learning:", input)

	// Conditionals
	/* Conditionals are essential in programming
	for making decisions and controlling the flow of your code.
	*/

	// BASIC
	// if statement

	num := 25
	if num > 18 {
		fmt.Println("You can vote")
	} // This is a basic demonstration how do we use if statement, Also we have else. Let's see


	age int := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")
	input, _ := age.ReadString('\n')
	fmt.Println("Your age is: ", input)

	if age > 18 {
		fmt.Println("You can travel via plane")
	} else {
		fmt.Println("Sorry, you can't travel via plane")
	}
}
