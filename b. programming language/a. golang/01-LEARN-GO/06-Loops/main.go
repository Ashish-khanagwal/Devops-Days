package main

import "fmt"

func main() {
	fmt.Println("Let's learn about using Loops in golang")

	// For Loop
	for i := 0; i <= 5; i++ {
		fmt.Println("iteration", i)
	}

	// Infinite loop
	for {
		fmt.Println("Infinite Loop")
		break // This will break the infinite
	}

	// iterating over slice
	num := [5]int{1, 2, 3, 4, 5}
	for index, value := range num {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}
}
