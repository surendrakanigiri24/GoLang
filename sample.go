package main

import "fmt"

func main() {
	// Static variable declaration
	conferenceName := "Go Conference" // We can call it as syntatic sugar it's same as (var conferenceName string = "value" ).
	const conferenceTickets = 50
	var remainingTickets uint = 50 // remaining ticktes could not be negative (Here uint is whole numbers)
	// var bookings [50]string        // In Go, array length should be fixed and it store only one data type values in one array
	var bookings = []string{} // It's a slice. For slice we need not to mention size

	// Messages
	// fmt.Println("Welcome to the world of booking tickets for", conferenceName)
	fmt.Printf("Welcome to the world of booking tickets for %v\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	// Dynamic Variable declaration w/o value
	var userName string
	var userTickets uint
	var email string

	fmt.Print("Enter your name: ")
	fmt.Scan(&userName) // Dynamically taking userName. Here & represents pointer(&username is address of variable userName to store input user value)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	// userTickets = 2
	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	// Logic
	remainingTickets = remainingTickets - userTickets

	// Array messages
	// bookings[0] = userName
	bookings = append(bookings, userName) // For slice we use append

	fmt.Printf("Whole Array/Slice is %v\n", bookings)
	fmt.Printf("First Array/Slice value is %v\n", bookings[0])
	fmt.Printf("Type of Array/Slice is %T\n", bookings)
	fmt.Printf("Length of Array/Slice is %v\n", len(bookings))

	fmt.Printf("Thank you %v for booking %v tikcets. You will  receive a confirmation email at %v in a short time \n", userName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
