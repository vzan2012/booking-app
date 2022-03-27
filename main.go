package main

import "fmt"

func main( ) {
	var conferenceName string = "Go Lang Conference"
	const conferenceTickets int = 50;
	var remainingTickets uint = 50;
	bookings := [] string{}

	fmt.Printf("Welcome to Our %v - Booking Application \n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available !!! \n" ,conferenceTickets,remainingTickets)
	fmt.Println("Get tickets to attend the meeting")

	// Data Types 
	var firstName string;
	var lastName string;
	var userTickets uint;
	var emailAddress string

	fmt.Println("Enter your First Name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name:")
	fmt.Scan(&lastName)
	fmt.Println("Enter number of tickets required:")
	fmt.Scan(&userTickets)
	fmt.Println("Enter your Email Address:")
	fmt.Scan(&emailAddress)

	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Bookings: %v \n", bookings[0])
	fmt.Printf("Bookings (Type): %T \n", bookings)
	fmt.Printf("Bookings (Length): %v \n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email to %v \n", firstName,lastName, userTickets, emailAddress);
	fmt.Printf("%v remaining tickets for %v", remainingTickets, conferenceName)
}