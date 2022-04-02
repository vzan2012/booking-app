package main

import (
	"fmt"
	"strings"
)

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
	for {
		fmt.Println("Enter your First Name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last Name:")
		fmt.Scan(&lastName)
		fmt.Println("Enter number of tickets required:")
		fmt.Scan(&userTickets)
		fmt.Println("Enter your Email Address:")
		fmt.Scan(&emailAddress)

		if userTickets <= remainingTickets {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)
	
			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking);
				firstNames = append(firstNames, names[0])
			}
	
			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email to %v \n", firstName,lastName, userTickets, emailAddress);
			fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
	
			fmt.Printf("The first name of bookings are: %v \n", firstNames)
	
			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("All Tickets are sold out !!!")
				break
			} 
		} else {
			fmt.Printf("We have only %v tickets, so you can't book %v tickets !!! \n", remainingTickets, userTickets)
			continue;	
		}
	}
}