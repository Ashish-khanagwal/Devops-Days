package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var username = "Ashish"
	fmt.Println("username: ",username)
	fmt.Printf("The data type of username is: %T \n", username)

	var myname = "Ashish"
	fmt.Println("Your name is: " + myname)
	fmt.Printf("Data type of myname is: %T \n", myname)
	
	// var age int
	// age = 30
	// fmt.Println("Your age is: " + strconv.Itoa(age))
	
	var age int = 30
	fmt.Println(age)
	fmt.Printf("Data type of age is: %T \n", age)

}