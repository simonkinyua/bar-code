package main

import (
	// "github.com/sanity-io/litter"
)


func bookTickets(firstName string, lastName string,email string, userTickets uint){

	remainingTickets = remainingTickets - userTickets
	userData := userData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}
	// bookings = append(bookings, firstName + "" + lastName)
	// userData["firstName"]=firstName
	// userData["lastName"]=lastName
	// userData["email"]=email
	// userData["ticketsBought"]=strconv.FormatUint(uint64(userTickets), 10)

	//litter.Dump(userData)

	bookings = append(bookings,userData)

}
