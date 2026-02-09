// Chapter 08: Structs
// This chapter introduces Structs (Structures), which are custom types
// that group different data types together under a single name.
package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"

// We define a custom 'User' struct. This is much cleaner than using maps
// because it allows mixed types (strings and uints) and fixed field names.
type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var bookings = make([]User, 0)

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
			fmt.Println("Invalid input data.")
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
		// Accessing struct fields using the dot '.' operator.
		firstNames = append(firstNames, booking.firstName)
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

	// Creating an instance of the 'User' struct.
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
