package main

import (
	"fmt"
	"strings"
)

// Package level variables
// Package level variables again no need to pass in functions because it’s accessible to all over the code
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{} // It's a slice. For slice we need not to mention size

func main() {

	// Greet user
	greetUser()

	for remainingTickets > 0 {

		// User input
		firstName, lastName, email, userTickets := getUserInput()

		// Validations
		isValidName, isValidEmail, isValidTicketNumber := isValidUserInput(firstName, lastName, email, userTickets, remainingTickets)
		if !isValidName || !isValidEmail || !isValidTicketNumber {
			fmt.Printf("Please enter valid input\n")
			continue
		}

		// Book tickets
		bookings, remainingTickets = bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

		// Print first names
		firstNames := getFirstNames(bookings)
		fmt.Printf("These are all our firstname of bookings %v\n", bookings)
		fmt.Printf("These are all our bookings %v\n", firstNames)

		var noTickets bool = remainingTickets == 0
		if noTickets {
			// End program
			fmt.Printf("Our conference is booked out. Come next year\n")
			break
		}
	}

}

// To greet users
func greetUser() {
	fmt.Printf("Welcome to the world of booking tickets for %v\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}

// Get User input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

// To print user firstnames
func getFirstNames(bookings []string) []string {
	// For each to have only first name
	firstNames := []string{}
	for _, booking := range bookings { // Here first value is index but we don't have any need of that. So, we are keeping _ instead of index. _ means in Go ignore that value.
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

// Validate user inputs
func isValidUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets < remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// Book tickets
func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, email string, conferenceName string) ([]string, uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v for booking %v tikcets. You will  receive a confirmation email at %v in a short time \n", firstName+""+lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings, remainingTickets

}
