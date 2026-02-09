// Chapter 03: Arrays and Slices
// This chapter introduces Slices, which are dynamic arrays in Go.
// Slices allow you to store multiple values of the same type in a single list.
package main

import "fmt"

func main() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"

	// Define a slice of strings. Slices are dynamic and can grow in size.
	// Syntax: []type{initial_values}
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)

	var firstName string
	var lastName string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	remainingTickets = remainingTickets - userTickets

	// append() adds a new element to the end of a slice.
	// It returns a new slice containing the additional element.
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	// Printing the entire slice shows all collected bookings.
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
