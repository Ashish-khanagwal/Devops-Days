package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's Learn Arrays")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("How's this course going on? \n")

	input, _ := reader.ReadString('\n')
	fmt.Println("It's pleasure to know that your course is going: ", input)

	// Now let's create an array
	// Syntax: var arrayName [size] dataType
	var array [5]int
	fmt.Println("The array is: ", array)

	// Let's initialize the array
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50
	fmt.Println("The array is: ", array)

	// Let's print the length of array
	fmt.Println("The lenght of the array is: ", len(array))

	// Another way of creating and initializing an array is
	names := [5]string{"Ashish", "Nancy", "Aneeket", "Muskan", "Nidhi"}
	fmt.Println("Name of you and your special people are: ", names)
	fmt.Println("Length of array names is: ", names)

	// EXAMPLE
	marks := [10]int{87, 56, 90, 99, 86, 54, 43, 33, 98, 97}
	fmt.Println("Length of marks array is: ", len(marks))

	// Accessing element in array
	fmt.Println("marks at 2nd place is: ", marks[2])

	// Iterating through array
	fruits := [4]string{"Apple", "Mango", "Grapes", "Watermelon"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Fruit at", i, "is", fruits[i])
	}
}
