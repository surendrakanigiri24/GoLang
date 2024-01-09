package main

import (
	"fmt"
	"sync"
	"time"
)

// Package level variables
// Package level variables again no need to pass in functions because it’s accessible to all over the code
const conferenceTickets = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// Structure define
type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// wait
var wg = sync.WaitGroup{}

// var bookings = []string{} // It's a slice. For slice we need not to mention size
// var bookings = []map[string]string{} // Map with empty slice
var bookings = []userData{}

func main() {

	// Greet user
	greetUser()

	// /for remainingTickets > 0 {

	// User input
	firstName, lastName, email, userTickets := getUserInput()

	// Validations
	isValidName, isValidEmail, isValidTicketNumber := isValidUserInput(firstName, lastName, email, userTickets, remainingTickets)
	if !isValidName || !isValidEmail || !isValidTicketNumber {
		fmt.Printf("Please enter valid input\n")
		return
		// continue
	}

	// Book tickets
	bookings, remainingTickets = bookTickets(remainingTickets, userTickets, firstName, lastName, email, conferenceName)

	// Send tickets
	wg.Add(1)
	go sendTicket(userTickets, firstName, lastName, email) // Here go keyword creates separate thread to execute in back ground

	// Print first names
	firstNames := getFirstNames()
	fmt.Printf("These are all our firstname of bookings %v\n", bookings)
	fmt.Printf("These are all our bookings %v\n", firstNames)

	var noTickets bool = remainingTickets == 0
	if noTickets {
		// End program
		fmt.Printf("Our conference is booked out. Come next year\n")
		// break
	}
	// }

	wg.Wait()

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
func getFirstNames() []string {
	// For each to have only first name
	firstNames := []string{}
	for _, booking := range bookings { // Here first value is index but we don't have any need of that. So, we are keeping _ instead of index. _ means in Go ignore that value.
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

// Book tickets
func bookTickets(remainingTickets uint, userTickets uint, firstName string, lastName string, email string, conferenceName string) ([]userData, uint) {
	remainingTickets = remainingTickets - userTickets

	// Structure
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Create a map of user
	// var userData = make(map[string]string) // Empty map declaration
	// userData["firstname"] = firstName
	// userData["lastname"] = lastName
	// userData["email"] = email
	// userData["ticktes"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %v for booking %v tikcets. You will  receive a confirmation email at %v in a short time \n", firstName+""+lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

	return bookings, remainingTickets

}

// Send ticket
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) // To delay 10 seconds
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
