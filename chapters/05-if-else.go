// Chapter 05: Conditionals
// This chapter teaches how to use if-else statements to validate user input
// and control the program flow based on conditions.
package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// Input Validation Logic
		// We check for name length, email format, and ticket availability.
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. Confirmation sent to %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// Logic to check if the conference is sold out.
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				// 'break' terminates the loop.
				break
			}
		} else {
			// Inform the user exactly what went wrong.
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}
			// 'continue' skips the rest of the current loop iteration.
			continue
		}
	}
}
