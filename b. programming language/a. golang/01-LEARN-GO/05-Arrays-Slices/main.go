package main

import "fmt"

func main() {
	fmt.Println("let's learn about Arrays and slices in golang")

	// Simple and basic way of declaring and storing values in arrays
	var numbers [5]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	numbers[3] = 4
	numbers[4] = 5
	// numbers[5] = 6 // It will show that index 5 is out of bound as array is of size 5
	fmt.Println("Arrays of numbers is:", numbers)

	// Another way of declare and store values in arrays
	var num = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array of num is:", num)

	// Storing strings in arrays
	names := [3]string{"Ashish", "Khanagwal", "Golang"}
	fmt.Println("Array of names is:", names)

	// Well we have one issue with arrays which can be overcome by using slices
	// Slices are dynamic in size where we do not have to define the size explicitly
	// It automatically expands the size when new element is added.

	// Here we are creating a slice of strings
	var namesSlice = []string{"Ashish", "Khanagwal", "Golang"}
	fmt.Println("Slice of names is:", namesSlice)

	// Here we are creating a slice of strings
	// we use make to create a new slice
	var names2 = make([]string, 3) // Type, length, (can specify capacity too)
	names2[0] = "Ashish"
	names2[1] = "Khanagwal"
	names2[2] = "Golang"
	fmt.Println("Slice of names is:", names2)
	fmt.Printf("length of Slice of names is:%v\n", len(names2))
	names2 = append(names2, "Khanagwal")
	fmt.Println("Slice of names is:", names2)
	fmt.Printf("length of Slice of names is:%v\n", len(names2))
}
