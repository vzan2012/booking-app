package main

import "fmt"

func main( ) {
	var conferenceName = "Go Lang Conference"
	const conferenceTickets = 50;
	var remainingTickets = 50;

	fmt.Printf("Welcome to Our %v - Booking Application \n",conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available !!! \n" ,conferenceTickets,remainingTickets)
	fmt.Println("Get tickets to attend the meeting")
}