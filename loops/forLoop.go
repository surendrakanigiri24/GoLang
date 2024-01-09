package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	var bookings = []string{} // It's a slice. For slice we need not to mention size

	fmt.Printf("Welcome to the world of booking tickets for %v\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
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

		// Logic
		if userTickets > remainingTickets {
			fmt.Printf("We have only %v tickets remaining\n", remainingTickets)
			continue
		}
		remainingTickets = remainingTickets - userTickets

		// Array messages
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v for booking %v tikcets. You will  receive a confirmation email at %v in a short time \n", firstName+""+lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// For each to have only first name
		firstNames := []string{}
		for _, booking := range bookings { // Here first value is index but we don't have any need of that. So, we are keeping _ instead of index. _ means in Go ignore that value.
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
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
