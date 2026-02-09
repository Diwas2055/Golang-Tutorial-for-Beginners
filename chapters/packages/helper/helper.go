// Package helper provides utility functions for the booking application.
// Functions in this package must start with a capital letter to be 'exported'
// (accessible from other packages).
package helper

import "strings"

// ValidateUserInput checks if the user provided valid data for a booking.
// It returns three booleans corresponding to name, email, and ticket validity.
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
