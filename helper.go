package main

import "strings"

// validateUserInput contains the core validation logic for booking data.
// It checks name lengths, email formatting, and ticket availability.
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	// Rule: Names must be at least 2 characters long
	isValidName := len(firstName) >= 2 && len(lastName) >= 2

	// Rule: Email must contain '@'
	isValidEmail := strings.Contains(email, "@")

	// Rule: Must book at least 1 ticket and not exceed availability
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
