package user

import (
	"fmt"
	"time"
)
func SendTicket(userBookedTickets uint, firstName string, lastName string, email string) {
	//Thanks Message
	fmt.Printf("Thanks for your booking, %v %v you booked %v tikcets, you will send you confirmation email to your registred email %v \n", firstName, lastName, userBookedTickets, email)
	//Delay the sending email for 5 seconds
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userBookedTickets, firstName, lastName)
	fmt.Println("###################")
	fmt.Printf("Sending ticket: \n %v to email address %v\n", ticket, email)
	fmt.Println("###################")
	// defer wg.Done() // matches the wg.Add(1) in main
}