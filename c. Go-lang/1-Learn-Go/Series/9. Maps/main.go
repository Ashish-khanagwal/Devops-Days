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
}
