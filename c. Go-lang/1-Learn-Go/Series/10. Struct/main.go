package main

import (
	"bufio"
	"fmt"
	"os"
)

// This is how we define a struct : To define a struct, use
// the type keyword followed by the struct name and a set of fields enclosed in curly braces.
type User struct { // One thing to be noticed that struct name should always start with uppercase letter
	FirstName, LastName string
	Age                 int
	Email               string
	Location            string
}

func main() {
	fmt.Println("Let's learn about Struct")

	Reader := bufio.NewReader(os.Stdin)
	fmt.Println("What we are gonna learn today? : ")
	input, _ := Reader.ReadString('\n')
	fmt.Println("We will be learning: ", input)

	user := User{"Ashish", "Khanagwal", 22, "ashish@gmail.com", "Delhi"}
	fmt.Println(user)
	fmt.Printf("User details are: %+v\n", user) // %+v will provide the full details about the struct
	fmt.Printf("User's name is %v and email is %v\n", user.FirstName, user.Email)

	// We can also define anonymous struct, see below
	// This is useful when we want to create a struct for one time use
	person := struct {
		Name           string
		age            int
		profession     string
		registrationId uint
	}{
		Name:           "Ashish",
		age:            22,
		profession:     "DevOps Engineer",
		registrationId: 032,
	}

	fmt.Println("This is how our anonymous struct will look like: ", person)
	fmt.Println(person.Name)
}
