package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Let's Learn about slices")

	// To create a slice, you don't need to specify the length upfront.
	numbers := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(numbers)
	fmt.Printf("Data type of numbers is: %T\n", numbers) // output => []int and this is slice

	// Slices can be sliced to extract subsets.
	// Syntax: slice[startIndexIncluding: endIndexExcluding]
	// Example:
	slice := numbers[1:3]
	fmt.Println("Sliced version of numbers is: ", slice)

	// Unlike Arrays, we use append to add elements in slice
	// Syntax: sliceName = append(sliceName, elementToAdd)
	// Example:
	numbers = append(numbers, 10)
	fmt.Println("The updated slice is: ", numbers)

	// Unexpected Behaviour
	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 976 // This will throw an error

	// To avoid this error, we can use append
	highScores = append(highScores, 555, 654, 976, 777)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
}
