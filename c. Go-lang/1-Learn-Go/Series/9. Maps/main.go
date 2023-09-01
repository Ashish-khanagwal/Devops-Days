package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's Learn about Maps in GoLang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Excited to learn about Maps in GoLang? Yes or No! ")
	input, _ := reader.ReadString('\n')
	fmt.Println("As you have answered the above question as: ", input)

	// This is how we create and initialize a map (as key-value pairs)
	languages := make(map[string]string)
	languages["GO"] = "GoLang"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"

	fmt.Println("The map is: ", languages)
	fmt.Println("Now let us see what happens when we try to access an element that does not exist in our map.")
	fmt.Println("The value of language at key 'C' is: ", languages["C"])

	// We can also create and initialize a map like this
	// Syntax: mapName := map[keyType]valueType{key1: value1, key2: value2, key3: value3}
	// Example:
	ages := map[string]int{"Ashish": 22, "Nidhi": 22, "Manish": 21, "Nancy": 20, "Muskan": 20}
	fmt.Println("People with ages are: ", ages)

	// Accessing the value from map
	fmt.Println("Age of Ashish is: ", ages["Ashish"])
	fmt.Println("GO is short for: ", languages["GO"])

	// Deleting from map
	delete(ages, "Manish")
	fmt.Println("Now the updated ages is: ", ages)

	// Modifiying the map
	languages["RB"] = "Ruby"
	fmt.Println("The updated languages are: ", languages)
}
