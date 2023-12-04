package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceTickets int = 50
	var conferenceName = "Go confrence year 2017"
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("Welcome to our %v,\n same as last year we have a total of %v tickets.\n", conferenceName, conferenceTickets)
	fmt.Println("\nWould you like to buy a ticket today? If so,")

	for {
		var firstName string
		var lastName string
		var userTickets uint
		var email string

		fmt.Println("please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("please enter your email.")
		fmt.Scan(&email)

		fmt.Println("please enter the amount of tickets you'd like to buy.")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for buying %v tickets.A confirmatin email will be sent to %v. \n\n", firstName, lastName, userTickets, email)
			fmt.Printf("Only %v tickets left now. \n", remainingTickets)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These are the first names of all users:%v \n\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("We are terribry sorry %v, we are sold out. Try again next year.\n", firstName)
				break
			}
		} else{
			fmt.Printf("We only have %v remaining tickets which means you cannot buy %v tickets.\n", remainingTickets, userTickets)
			continue
	}

	}

}
