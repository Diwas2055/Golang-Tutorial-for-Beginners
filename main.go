// Capstone: Booking Application
// This is the final project that integrates all concepts learned in the tutorial,
// including structs, functions, loops, validation, and concurrency.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Package-level constants and variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

// UserData groups all information about a single booking
type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// sync.WaitGroup is used to wait for all asynchronous tasks (sending emails) to finish
var wg = sync.WaitGroup{}

func main() {
	// Greet the user and show initial state
	greetUsers()

	for {
		// 1. Collect user information
		firstName, lastName, email, userTickets := getUserInput()

		// 2. Validate user input using logic in helper.go
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// 3. Update the booking records
			bookTicket(userTickets, firstName, lastName, email)

			// 4. Start an asynchronous task to "send" the ticket
			// We increment the WaitGroup counter before starting the goroutine.
			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			// 5. Display current bookings
			firstNames := getFirstNames()
			fmt.Printf("Current bookings (first names): %v\n", firstNames)

			// 6. Check if the conference is sold out
			if remainingTickets == 0 {
				fmt.Println("The conference is fully booked. See you next year!")
				break
			}
		} else {
			// Specific error messages for invalid input
			if !isValidName {
				fmt.Println("Error: First or last name is too short (min 2 chars).")
			}
			if !isValidEmail {
				fmt.Println("Error: Email address must contain an '@' symbol.")
			}
			if !isValidTicketNumber {
				fmt.Printf("Error: Invalid number of tickets. Only %v remaining.\n", remainingTickets)
			}
		}
	}

	// Wait for all background goroutines (email sending) to complete before exiting
	wg.Wait()
}

// greetUsers prints the application header
func greetUsers() {
	fmt.Printf("Welcome to the %v Booking Application\n", conferenceName)
	fmt.Printf("Total Tickets: %v | Available: %v\n", conferenceTickets, remainingTickets)
	fmt.Println("--------------------------------------------------")
}

// getFirstNames extracts a list of first names from the bookings slice
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// getUserInput prompts the user and collects data from stdin
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// bookTicket updates the remaining tickets and adds the user to the list
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("Success! %v %v booked %v tickets. Confirmation sent to %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Tickets remaining: %v\n", remainingTickets)
}

// sendTicket simulates a long-running process (like sending an email) using a goroutine
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	// Simulate a 10-second delay for email processing
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n##################################################")
	fmt.Printf("SIMULATED EMAIL: Sending ticket...\n%v\nto address %v\n", ticket, email)
	fmt.Println("##################################################")

	// Notify the WaitGroup that this task is complete
	wg.Done()
}
