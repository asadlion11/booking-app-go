package user

import "strings" // The strings package offers handy functions for string manipulation—e.g., Join, Split, Contains, Trim, and case conversions.

//Valiadate User Input Function
func ValidateUserInput(firstName string , lastName string, email string, userBookedTickets uint, remainingTickets uint) (bool, bool, bool) {
	//Checking user input is valid
	var isValidNames bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidUserBookedTickets bool = userBookedTickets > 0 && userBookedTickets <= remainingTickets
	//Return the values
	return isValidNames, isValidEmail, isValidUserBookedTickets
}