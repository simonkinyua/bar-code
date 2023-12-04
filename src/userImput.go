package main

import(
	"fmt"
	// "strings"
)

func getUserImput() (string, string, uint, string) {
	greetings();  // get and validate user imput

	var firstName string
	var lastName string
	var userTickets uint
	var email string

	
	fmt.Println("please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("\nPlease enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("please enter your email.")
	fmt.Scan(&email)
	fmt.Println("please enter the amount of tickets you'd like to buy.")
	fmt.Scan(&userTickets)

		return firstName, lastName, userTickets, email
    
}
		
