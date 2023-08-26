package main

import "fmt"

func main() {
	fmt.Println("Let's begin with pointers")

	var ptr *int
	fmt.Println("The value of pointer is ", ptr)
	// It will return value as <nil> because we haven't initialized it without any value
	// But you've learnt how we create pointer in golang
	// so let us initialize the variable now and check again!

	num := 32
	ptr = &num

	// here "&" it is referencing the address and ptr holds the value of address

	fmt.Println("the value of pointer is ", ptr)  // This will print out the address of num
	fmt.Println("the value of pointer is ", *ptr) // This will print out the value of num

	// We can also modify the value of pointer, have a look how
	*ptr = 45
	fmt.Println("The updated value of pointer is ", *ptr) // Value is updated as 45

	// We also can perform mathematical calulation to our pointer, have a look how
	*ptr *= 2
	fmt.Println("The calculated value of pointer is ", *ptr)
}
