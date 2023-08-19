package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your professin: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your response ", input)
}
