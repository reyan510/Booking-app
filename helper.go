package main

import "strings"

//export a function by captalizing first word of function
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	// A Go function can return multiple values
	return isValidName, isValidEmail, isValidTicketNumber
}
