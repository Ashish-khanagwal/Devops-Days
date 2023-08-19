package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza shop")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Thank you for ordering Pizza, Please give our pizza a rating from 1-5")
	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating our pizza, you rated: ", input)

	// numRating = input + 1
	// Look at the above code, you'll find out that "input" is of string type and "1" is integer so, It can't be added to string.
	// We will use the typeconversion here, and the package used to achieve the same is "strconv"

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1 is Added to your rating: ", numRating+1)
	}
}
