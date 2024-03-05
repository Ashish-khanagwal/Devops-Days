package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "GO Conference"
var conferenceTickets int = 50
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

var wg = sync.WaitGroup{}

type UserData struct {
	firstName     string
	lastName      string
	email         string
	bookedTickets uint
}

func main() {

	greetUser()

	firstName, lastName, email, bookedTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, bookedTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTickets(bookedTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTickets(bookedTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("First name of the users booked event tickets are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Sorry, we are out of tickets for the event. Come back next year 🙃.")
			// break
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
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to the %v ticket booking application\n", conferenceName)
	fmt.Printf("We have %v tickets for the event. And %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend the conference")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var bookedTickets uint

	fmt.Print("Enter your first-name, and last-name: ")
	fmt.Scanln(&firstName, &lastName)

	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	fmt.Print("How many tickets do you want to book? ")
	fmt.Scanln(&bookedTickets)

	return firstName, lastName, email, bookedTickets
}

func bookTickets(bookedTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - bookedTickets

	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		bookedTickets: bookedTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets for the event. Confirmation regarding your booking will be sent via email at your %v id\n", firstName, lastName, bookedTickets, email)
	fmt.Printf("Now we have %v tickets remaining for the %v.\n", remainingTickets, conferenceName)

}

func sendTickets(bookedTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", bookedTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending tickets: \n %v to email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
