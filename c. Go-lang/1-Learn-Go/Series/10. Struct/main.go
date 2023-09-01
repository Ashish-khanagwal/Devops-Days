package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Let's learn about Struct")

	Reader := bufio.NewReader(os.Stdin)
	fmt.Println("What we are gonna learn today? : ")
	input, _ := Reader.ReadString('\n')
	fmt.Println("We will be learning: ", input)
}
