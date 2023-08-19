package main

import "fmt"

func main() {
	fmt.Println("Variables")

	var userName = "Ashish"
	fmt.Println(userName)
	fmt.Printf("Data type of username is: %T \n", userName)

	userName2 := "Ashish"
	fmt.Println(userName2)
	fmt.Printf("Data type of username2 is: %T \n", userName2)

	var rollNo uint8 = 255
	fmt.Println(rollNo)
	fmt.Printf("Data type of rollNo is: %T \n", rollNo)

	var marks float32 = 95.869789
	fmt.Println(marks)
	fmt.Printf("Data type of marks is: %T \n", marks)

	var isPass bool = true
	fmt.Println(isPass)
	fmt.Printf("Data type of isPass is: %T \n", isPass)

}
