package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, bookedTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicketNumber := bookedTickets > 0 && bookedTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
