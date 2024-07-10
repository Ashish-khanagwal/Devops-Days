package main

import "fmt"

func main() {
	fmt.Println("Let's learn about maps")

	// A basic map in Go is a collection of key-value pairs where each key is unique.
	// The syntax for a map is map[keyType]valueType
	// Here, we are creating a map with key type as string and value type as int
	age := map[string]int{
		"Ashish":   23,
		"Priyanka": 20,
	}
	fmt.Println("Ashish's age", age["Ashish"])
	fmt.Println("Priyanka's age", age["Priyanka"])

	// Another exmple of maps
	var scores = make(map[string]int) // This will create an empty map

	scores["Ashish"] = 100 // Adding key-value pair to the map
	scores["Priyanka"] = 90
	fmt.Println("Ashish's score", scores["Ashish"])
	fmt.Println("Priyanka's score", scores["Priyanka"])

	scores["Navya"] = 85
	fmt.Println("Navya's score", scores["Navya"])
	fmt.Println(scores)

	scores["Ashish"] = 99 // Modifying the value of a key
	fmt.Println("Ashish's score", scores["Ashish"])
	fmt.Println(scores)

	delete(scores, "Navya") // Deleting a key-value pair from the map
	fmt.Println(scores)
}
