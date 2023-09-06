package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main() {
	fmt.Println("Now, different methods of taking inputs")

	// Method 1
	var num int
	fmt.Println("Enter a number: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("you entered: ", num)

	// Method 2
	// If we have to take multiple inputs then we can simple do this

}
