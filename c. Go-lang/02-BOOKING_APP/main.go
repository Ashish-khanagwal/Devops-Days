package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "GO Conference"
	var conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to the %v ticket booking application\n", conferenceName)
	fmt.Printf("We have %v tickets for the event. And %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets from here to attend the conference")

	for {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicketNumber := bookedTickets > 0 && bookedTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - bookedTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets for the event. Confirmation regarding your booking will be sent via email at your %v id\n", firstName, lastName, bookedTickets, email)
			fmt.Printf("Now we have %v tickets remaining for the %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("First name of the users booked event tickets are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Sorry, we are out of tickets for the event. Come back next year 🙃.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("email is invali, you didn't entered '@'")
			}
			if !isValidTicketNumber {
				fmt.Println("number of ticket entered is invalid")
			}
		}
	}
}
