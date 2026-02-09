// Chapter 02: Variables and User Input
// This chapter covers how to store data using variables and constants,
// and how to accept input from the user via the terminal.
package main

import "fmt"

func main() {
	// 'const' defines a value that cannot be changed after declaration.
	const conferenceTickets int = 50

	// 'var' defines a variable with an explicit type.
	// 'uint' is an unsigned integer (always positive).
	var remainingTickets uint = 50

	// ':=' is the short declaration syntax.
	// It infers the type (string) and is only available inside functions.
	conferenceName := "Go Conference"

	// fmt.Printf allows formatted strings. %v is the default format for values.
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	// Declare variables to hold user input
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask for user input
	fmt.Println("Enter Your First Name: ")
	// fmt.Scanln reads input from the console.
	// The '&' symbol provides the memory address (pointer) of the variable.
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	// Basic arithmetic to update remaining tickets
	remainingTickets = remainingTickets - userTickets

	// Summary output
	fmt.Printf("Thank you %v %v for booking %v tickets. Confirmation sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
