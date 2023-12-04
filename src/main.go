package main

import (
	"fmt"
	"strings"
	"sync"
)

var remainingTickets uint = 50
var bookings = make([] userData,0)


type userData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	    for remainingTickets > 0 && len(bookings) < 50{ 
		firstName, lastName, userTickets, email := getUserImput()
		isValidEmail, isValidName, isValidUserTickets := validateUserInfo(firstName, lastName,email, userTickets)

		if isValidEmail && isValidName && isValidUserTickets{
 
			bookTickets(firstName, lastName,email, userTickets)

			wg.Add(1)
			go sendTicket(firstName, lastName, userTickets, email)


			firstNames := getFirstNames()
			fmt.Printf("Thank you user %v %v for buying %v tickets today.\n The confirmation email have be sent to %v.\n",firstName, lastName, userTickets, email)

	        fmt.Printf("\nThese are the first names of all users:\t %+v \n\n", firstNames)
			fmt.Printf("\n%v remaining tickets only.\n", remainingTickets)

			fmt.Printf("\nUser Data is as follows: %+v\n",bookings)

		if remainingTickets==0{
			fmt.Println("We are terribly sorry for we are booked out, try again next year.")
		}

		} else{ 

		if !isValidName {
			fmt.Println("The first or last name you entered is too short or invalid")
			}
			if !isValidEmail {
				fmt.Println("The email you provided is invalid")
			}
			if !isValidUserTickets {
				fmt.Println("The number of tickets is ivalid")
			}
		}

	}
	wg.Wait()
}

func validateUserInfo(firstname string, lastname string, Uemail string, usertickets uint) (bool, bool, bool){
	isValidName := len(firstname) >= 2 && len(lastname) >= 2
	isValidEmail := strings.Contains(Uemail, "@")
	isValidUserTickets := remainingTickets > 0 && usertickets <= remainingTickets

	return isValidEmail, isValidName, isValidUserTickets
}
	    
func getFirstNames ()[] string{
	var firstNames =[]string{}

	for _, booking := range bookings {
		firstNames = append(firstNames,booking.firstName)
	}

	return firstNames
}