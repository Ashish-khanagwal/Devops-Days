package main

import "fmt"

func main() {
	fmt.Println("Now we will learn about Switch statements")

	// Basic Switch Statement
	var day string
	fmt.Print("Enter any day of the week: ")
	fmt.Scanln(&day)
	switch day {
	case "Monday", "monday":
		fmt.Println("Today is Monday")
	case "Tuesday", "tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday", "wednesday":
		fmt.Println("Today is Wednesday")
	case "Thursday", "thursday":
		fmt.Println("Today is Thursday")
	case "Friday", "friday":
		fmt.Println("Today is Friday")
	case "Saturday", "saturday":
		fmt.Println("Today is Saturday")
	case "Sunday", "sunday":
		fmt.Println("Today is Sunday")
	default:
		fmt.Println("Invalid day")
	}
	// In Go, each case in a switch statement automatically breaks at the end

	// Fallthrough just let you go to the next case without breaking
	chocolate := "caramel"
	switch chocolate {
	case "caramel":
		fmt.Println("I love caramel chocolate")
		fallthrough
	case "dark":
		fmt.Println("I love dark chocolate")
	}

	// We can also use expressions in switch cases
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	switch {
	case age >= 18:
		fmt.Println("You can vote")
	case age < 18:
		fmt.Println("You can't vote")
	default:
		fmt.Println("Invalid age")
	}
}
