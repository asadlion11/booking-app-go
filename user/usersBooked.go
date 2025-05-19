package user

import (
	"fmt"
	//"strconv"
)

//Store User Data using struct
type UserInfo struct {
	firstName string
	lastName string
	email string
	userBookedTickets uint
}

//Display list of users who booked tickets
func DisplayUsersBooked(usersBooked []UserInfo, firstName , lastName , email string, userBookedTickets uint) []UserInfo{

	//Users data into map then into slice using map
	// userData := map[string] string {
	// 	"firstName": firstName,
	// 	"lastName": lastName,
	// 	"email": email,
	// 	"userBookedTickets": strconv.FormatUint(uint64(userBookedTickets), 10),
	// }

	var userData = UserInfo {
		firstName: firstName,
		lastName: lastName,
		email: email,
		userBookedTickets: userBookedTickets,
	}

	
	//add the new user into the list of users who booked tickets
	usersBooked = append(usersBooked, userData)

	//Diplsay list of users who booked tickets
	//"[]": literal list brackets; use "%s" for strings or "%d" for ints, e.g. "[%s]" or "[%d]".
	fmt.Printf("List of users who booked tickets are (%d total): \n", len(usersBooked))

	//Print each booked user(firts name and last name) with a counter(i+1 and valur u = usersBooked[current i no as index])
	 // Print all bookings
	for i, user := range usersBooked {
		fmt.Printf("%d. Full Name: %s %s, Email: %s, Booked: %d tickets\n",
			i+1,
			user.firstName,   // Use struct fields from the loop variable
			user.lastName,
			user.email,
			user.userBookedTickets,
		)
	}
    fmt.Println()

	return usersBooked
}