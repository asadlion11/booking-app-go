package user

import "fmt"

func WelcomeUser(conferenceName string, totalTickets int, remainingTickets uint) {
	//Welcome Message
	fmt.Printf("Welcome to our %v Booking App.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", totalTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend the conference.")
}