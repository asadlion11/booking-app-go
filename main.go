package main

import (
	"fmt" //IO package
	//"strings" // The strings package offers handy functions for string manipulation—e.g., Join, Split, Contains, Trim, and case conversions.
)

func main() {

	//The app can use only one user at a time
	// var conferenceName string = "Go Conference"
	// const totalTickets int = 50 
	// var remainingTickets uint = 50 //unit accepts postive numbers only

	// //Welcome message usinf Printf(Fotmatted output): %v is a placeholder for any value
	// fmt.Printf("Welcome to our %v Booking App.\n", conferenceName)
	// fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	// fmt.Println("Get your ticket here to attend the conference.")

	// // //Welcome message using Println(New line)
	// // fmt.Println("Welcome to our" ,conferenceName, "Booking App.")
	// // fmt.Println("We have a total of ", confereceTickets, "tickets and", remainingTickets, "are still available.")
	// // // fmt.Print("Get your ticket here to attend the conference.")
	// // fmt.Println("Get your ticket here to attend the conference.") //new line

	// //Poniter 
	// // var myName string = "Anas"
	// // fmt.Println(myName)
	// // fmt.Println(&myName) //memory address of the variable

	// //store list of users who booked tickets
	// //var bookings[50] string

	// //store list of users who booked slice
	// var usersBooked[] string
	

	// var firstName string
	// var lastName string
	// var email string
	// var userBookedTickets uint
	// //User Input
	// fmt.Print("Enter your first name: ")
	// fmt.Scan(&firstName)

	// fmt.Print("Enter your last name: ")
	// fmt.Scan(&lastName)

	// fmt.Print("Enter your email: ")
	// fmt.Scan(&email)

	// fmt.Print("Enter number of tickets you want to book: ")
	// fmt.Scan(&userBookedTickets)

	// //add the user into bookings array store(list)
	// //bookings[0] = firstName + " " + lastName

	// //add the user into bookings array store(list)
	// usersBooked = append(usersBooked, firstName + " " + lastName)

	// //Display message to the user
	// fmt.Printf("Thanks for your booking, %v %v you booked %v tikcets, you will send you confirmation email to your registred email %v \n", firstName, lastName, userBookedTickets, email)

	// //Remaining Tickets Logic 
	// remainingTickets = remainingTickets - userBookedTickets
	// fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)

	// //Diplsay list of users who booked tickets we used array
	// //fmt.Printf("List of users who booked tickets\n %v\n", bookings)
	// // fmt.Printf("Array type: %T \n", bookings)
	// // fmt.Printf("Array Length: %v \n", len(bookings))

	// //Diplsay list of users who booked tickets we used slice
	// fmt.Printf("List of users who booked tickets: %v \n", usersBooked)
	// // fmt.Printf("First value is slice: %v \n", usersBooked[0])


	//The app can use multiple users at a time(we use loop: for loor infinite loop)
	var conferenceName string = "Go Conference"
	const totalTickets int = 50 
	var remainingTickets uint = 50 

	//Welcome message
	fmt.Printf("Welcome to our %v Booking App.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend the conference.")

	//list of users who booked the tickets
	var usersBooked[] string
	
	for {
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

		//add the new user into the list of users who booked tickets
		usersBooked = append(usersBooked, firstName + " " + lastName)
		
		//Display message to the new user(feedback)
		fmt.Printf("Thanks for your booking, %v %v you booked %v tikcets, you will send you confirmation email to your registred email %v \n", firstName, lastName, userBookedTickets, email)

		//Update the remaining tickets
		remainingTickets = remainingTickets - userBookedTickets
		fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)

		//Diplsay list of users who booked tickets
		//"[]": literal list brackets; use "%s" for strings or "%d" for ints, e.g. "[%s]" or "[%d]".
		fmt.Printf("List of users who booked tickets are %d and they are: \n", len(usersBooked))

		//Print each booked user(firts name and last name) with a counter(i+1 and valur u = usersBooked[current i no as index])
		for i, u := range usersBooked  { 
			// range usersBooked: walks the slice from start to end. 
			// i: starts at 0
			// u: is the value at that index, i.e. usersBooked[i]
			fmt.Printf("%d. %s\n", i+1, u)
		}

		//stope the infinite loop if the remaining tickets are 0(if and break)
		if remainingTickets == 0 {
			fmt.Println("All the tickets are booked, get back the next year")
			//stope the infinite loop
			break
		}
		

		//Print the first names of booked users
		//we use for-each loop
	// 	firstNames := []string{}
	// 	for _, bookingUser := range usersBooked{
	// 		var names = strings.Fields(bookingUser)
	// 		firstNames = append(firstNames, names[0])
	// 	}

	// 	fmt.Println("First names of booked users are:")
	// 	fmt.Printf("%s \n", strings.Join(firstNames, "\n"))
	}
}