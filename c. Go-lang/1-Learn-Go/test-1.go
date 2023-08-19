package main

import "fmt"

func main() {
  fmt.Println("Starting Textion Server")

  userName := "Ashish"
  userAge := 22

  fmt.Println("My name is", userName, "and my age is", userAge)
  fmt.Printf("Data type of name is %T and data type of age is %T", userName, userAge)
  fmt.Println()
  fmt.Printf("Hello, %v how are you?", userName)
  fmt.Println() 
  var name string
  fmt.Scan(&name)
  fmt.Printf("Hello %v, welcome to the milkyway galaxy", name)
}
