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

	// Welcome User 
	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for {
		// User Input
		firstName, lastName, emailAddress, userTickets := getUserInput() 

		// Validation 
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, emailAddress, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			// Book Conference Tickets 
			bookTickets(remainingTickets, userTickets, bookings, firstName, lastName, emailAddress, conferenceName)
	
			firstNames := getFirstNames(bookings)
			fmt.Printf("The first name of bookings are: %v \n", firstNames)
					
			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("All Tickets are sold out !!!")
				break
			} 
		} else {
			if !isValidName {
				fmt.Println("First Name or Last Name is invalid !!!")
			}
			if !isValidEmail {
				fmt.Println("The given Email address doesn't contain '@' symbol - invalid !!!")
			}			
			if !isValidTicketNumber {
				fmt.Println("Number of tickets given is invalid !!!")
			}
			// fmt.Printf("We have only %v tickets, so you can't book %v tickets !!! \n", remainingTickets, userTickets)
			// continue;	
		}
	}
}

func greetUser(confName string, confTickets int, remTickets uint) {
	fmt.Println("Welcome to the Conference !!!")	
	fmt.Printf("Welcome to Our %v - Booking Application \n",confName)
	fmt.Printf("We have total of %v tickets and %v are still available !!! \n" ,confTickets,remTickets)
	fmt.Println("Get tickets to attend the meeting")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking);
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, emailAddress string, userTickets uint, remainingTickets uint) (bool,bool,bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail:= strings.Contains(emailAddress, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets 
	
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, emailAddress, userTickets 
}

func bookTickets(remainingTickets uint, userTickets uint, bookings []string, firstName string, lastName string, emailAddress string, conferenceName string) {
	remainingTickets -= userTickets
	bookings = append(bookings, firstName+" "+lastName)


	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email to %v \n", firstName,lastName, userTickets, emailAddress);
	fmt.Printf("%v remaining tickets for %v \n", remainingTickets, conferenceName)
}