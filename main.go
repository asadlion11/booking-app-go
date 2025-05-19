package main

import (
	"booking-app/user"
	"fmt" //IO package
	"sync"
)
var conferenceName string = "Go Conference"
const totalTickets int = 50 
var remainingTickets uint = 50
// var usersBooked[] map[string] string
var usersBooked []user.UserInfo
var wg = sync.WaitGroup{}

//The app can use multiple users at a time(we use loop: for loor infinite loop)
func main() {
	//Greet User
	user.WelcomeUser(conferenceName, totalTickets, remainingTickets)
		//Get User Input
		firstName, lastName, email, userBookedTickets := user.GetUserInput()

		//Validate User Input
		isValidNames, isValidEmail, isValidUserBookedTickets := user.ValidateUserInput(firstName, lastName, email, userBookedTickets, remainingTickets)

		//Check Validation
		if  isValidNames && isValidEmail && isValidUserBookedTickets{
			//Update the remaining tickets / booking logic
			remainingTickets = remainingTickets - userBookedTickets
			fmt.Printf("%v tickets are remaining for %v \n", remainingTickets, conferenceName)

			//Send Ticket //go: concurrency allow use multiple use at the same time
			//Add(number of threads thet the main thread should wait)
		  // 1) Schedule SendTicket in a goroutine...
		  //Th work is going in the backgorund
			wg.Add(1)
			go func() {
				//Inside that goroutine you defer wg.Done() so it signals completion when SendTicket returns.
				//is that defer wg.Done() registers the Done() call immediately, guaranteeing it will run after SendTicket returns—no matter how SendTicket exits (even if it panics).
				//Using defer at the top of the goroutine is the idiomatic way to make sure every Add(1) always has a matching Done(), even in error cases.
				defer wg.Done()
				user.SendTicket(userBookedTickets, firstName, lastName, email)
			}()

			//Call displayUsersBooked function to display the list of users who booked tickets
			usersBooked = user.DisplayUsersBooked(usersBooked, firstName, lastName, email, userBookedTickets)


			//stop the infinite loop if the remaining tickets are 0(if and break)
			if remainingTickets == 0 {
				fmt.Println("All the tickets are booked, get back the next year")
				//stope the infinite loop
				//break
			}
		} else {
			if !isValidNames {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email you entered does't contain @ sign")
			}
			if !isValidUserBookedTickets {
				fmt.Println("Numbers if tickets you entered is invalid")
			} 	
		}
		//waits all the threads that were added thet mention the Add function to be done, doing the application can exit
		wg.Wait()
}

