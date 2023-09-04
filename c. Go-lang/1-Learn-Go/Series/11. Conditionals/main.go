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

	var age int
	fmt.Println("Enter your age: ")
	_, err := fmt.Scan(&age)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else if age >= 18 {
		fmt.Println("Your age is: ", age)
		fmt.Println("You can travel via plane")
	} else {
		fmt.Println("Sorry, you can't travel via plane")
	}

	var height, weight int
	fmt.Println("Enter your height (in cm) and weight (in kg): ")
	_, err = fmt.Scan(&height, &weight)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("your height and weight is: ", height, weight)
}
