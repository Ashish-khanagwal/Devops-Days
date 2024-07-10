package main

import "fmt"

func main() {
	fmt.Println("Now we will learn about conditionals")

	// Simple conditionals
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}

	// some additional condition with new example
	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)
	if age >= 18 {
		fmt.Println("you can vote")
	} else {
		fmt.Println("you cannot vote")
	}

	// More complex
	fmt.Print("Enter your age: ")
	var yourAge int
	fmt.Scanln(&yourAge)

	if yourAge >= 18 {
		if yourAge >= 22 {
			fmt.Println("you can watch Animal movie")
		} else {
			fmt.Println("Padhai likhai karo IAS IPS bano")
		}
	} else if (yourAge < 18) && (yourAge >= 15) {
		fmt.Println("you can watch Comedy movie")
	} else {
		fmt.Println("Go and watch chota bheem on pogo")
	}
}
