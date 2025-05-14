package main
import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const totalTickets int = 50 
	var remainingTickets uint = 50 //unit accepts postive numbers only

	//Welcome message usinf Printf(Fotmatted output): %v is a placeholder for any value
	fmt.Printf("Welcome to our %v Booking App.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend the conference.")

	// //Welcome message using Println(New line)
	// fmt.Println("Welcome to our" ,conferenceName, "Booking App.")
	// fmt.Println("We have a total of ", confereceTickets, "tickets and", remainingTickets, "are still available.")
	// // fmt.Print("Get your ticket here to attend the conference.")
	// fmt.Println("Get your ticket here to attend the conference.") //new line

	//Poniter 
	// var myName string = "Anas"
	// fmt.Println(myName)
	// fmt.Println(&myName) //memory address of the variable

	var firstName string
	var lastName string
	var email string
	var userBookedTickets uint
	//User Input
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets you want to book: ")
	fmt.Scan(&userBookedTickets)
	 
	//Display message to the user
	fmt.Printf("Thanks for your booking, %v %v you booked %v tikcets, you will send you confirmation email to your registred email %v \n", firstName, lastName, userBookedTickets, email)

	//Remaining Tickets Logic 
	remainingTickets = remainingTickets - userBookedTickets
	fmt.Printf("%v tickets are remaining for %v", remainingTickets, conferenceName)
}