// Chapter 04: Loops
// This chapter demonstrates how to repeat code execution using loops.
// Go only has one loop keyword: 'for'.
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

	// An infinite loop starts with 'for' followed by a block.
	// This allows the program to keep running until explicitly stopped.
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

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. Confirmation sent to %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// Iterate over the slice to extract only first names.
		// 'range' returns the index and the value of each element.
		// We use '_' to ignore the index.
		firstNames := []string{}
		for _, booking := range bookings {
			// strings.Fields splits a string by whitespace.
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of current bookings: %v\n", firstNames)
	}
}
