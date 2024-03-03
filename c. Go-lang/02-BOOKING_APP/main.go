package main

import "fmt"

func main() {
	conferenceName := "GO Conference"
	var conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v ticket booking application\n", conferenceName)
	fmt.Printf("We have %v tickets for the event. And %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend the conference")

	var firstName string
	var lastName string
	var email string
	var age int
	var bookedTickets uint

	fmt.Print("Enter your first-name, and last-name: ")
	fmt.Scanln(&firstName, &lastName)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("How many tickets do you want to book? ")
	fmt.Scanln(&bookedTickets)

	remainingTickets = remainingTickets - bookedTickets

	fmt.Printf("Thank you %v %v for booking %v tickets for the event. Confirmation regarding your booking will be sent via email at your %v id\n", firstName, lastName, bookedTickets, email)
	fmt.Printf("Now we have %v tickets remaining for the %v.\n", remainingTickets, conferenceName)
}
