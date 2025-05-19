package user

import "fmt"

//Get user input function
func GetUserInput() (string , string, string, uint){
	//store user info
	var firstName string
	var lastName string
	var email string
	var userBookedTickets uint

	//Get the new user's info(the user who is booking the ticket)
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userBookedTickets)

	return firstName, lastName, email, userBookedTickets
}