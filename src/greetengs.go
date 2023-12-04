package main

import(
	"fmt"
)

func greetings(){
	const conferenceTickets int = 50
	conferenceName := "Go confrence year 2017"

	fmt.Printf("\nWelcome to our %v,\n same as last year we have a total of %v tickets.\n", conferenceName, conferenceTickets)
	fmt.Println("\nWould you like to buy a ticket today? If so,")

}