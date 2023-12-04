package main

import (
	"fmt"
	"time"
)

	func sendTicket(firstName string, lastName string,userTickets uint, email string){


		ticket := fmt.Sprintf("%v user tickets for: %v %v", userTickets, firstName, lastName)
		fmt.Println("*************")
		fmt.Printf("\n Sending %v \n to email adress %v.\n\n", ticket, email)
		fmt.Println("*************")
		time.Sleep(10 * time.Second)

		wg.Done()

}