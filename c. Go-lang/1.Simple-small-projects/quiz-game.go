package main

import (
  "bufio"
  "fmt"
  "strings"
  "os"
)

func main() {
  fmt.Println("Welcome to the quiz game")

  fmt.Printf("Enter your name: ")
  var name string
  fmt.Scan(&name)
  fmt.Printf("Hello %v, Welcome to the game! \n", name)

  fmt.Printf("Enter your age: ")
  var age uint
  fmt.Scan(&age)

  if age >= 10 {
    fmt.Printf("You're %v years old, Yay! you can play the game\n", age)
  } else {
    fmt.Printf("Sorry! you can't play the game")
    return
  }
    
  fmt.Printf("Who was Indian Cricket Team's captain? - 'Virat Kohli' or 'M.S Dhoni'\n ")
  scanner := bufio.NewScanner(os.Stdin) // This will read the whole line instead of reading individual characters.
  fmt.Printf("Answer the above question: ")

  if scanner.Scan() {
    answer := scanner.Text()
    answer = strings.ToLower(answer)
    
    if answer == "virat kohli" {
      fmt.Println("congratulations! you answered correctly and you won a chance to click a selfie with virat kohli")      
    } else if answer == "m.s dhoni" {
      fmt.Println("Oops! you're answer is wrong. you looses")
    } else {
      fmt.Println("please, read the question carefully")
    }
  }
}
