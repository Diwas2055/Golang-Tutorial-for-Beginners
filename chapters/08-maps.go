// Chapter 07: Maps
// This chapter introduces Maps, which are key-value pairs used to store
// more complex and structured data than slices alone.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"

// We define a slice of maps. Each map represents a single user booking.
var bookings = make([]map[string]string, 0)

func main() {
	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Conference sold out.")
				break
			}
		} else {
			fmt.Println("Invalid input, please try again.")
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		// Accessing values from a map using the key.
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint

	fmt.Println("Enter First Name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Enter Last Name: ")
	fmt.Scanln(&lastName)
	fmt.Println("Enter Email: ")
	fmt.Scanln(&email)
	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// Creating a map to store user data.
	// Syntax: make(map[keyType]valueType)
	var user = make(map[string]string)
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email

	// Maps can only have one type for values.
	// To store an integer as a string, we use strconv.FormatUint.
	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
}
