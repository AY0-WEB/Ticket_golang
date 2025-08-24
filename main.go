package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference"

const conferenceTickets = 50

var remainingTickets = 50
var Bookings = make([]userData, 0)

type userData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets int
}

var wg = sync.WaitGroup{}

func main() {
	for {
		greetUser()

		firstName, lastName, emailAddress, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNum := validDoc(firstName, lastName, emailAddress, userTickets)

		if isValidEmail && isValidName && isValidTicketNum {

			bookTicket(userTickets, firstName, lastName, emailAddress)
			wg.Add(1)
			go sendTicket(firstName, lastName, emailAddress, userTickets)

			firstNames := getFirstName()
			fmt.Printf("First names of those that booked %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("We are fully booked out.. pls come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your last name or first name is to short")
			}
			if !isValidEmail {
				fmt.Println("Email address doen't contain the @ sign")
			}
			if !isValidTicketNum {
				fmt.Println("invalid Ticket number")
			}
		}
		wg.Wait()
	}
}
func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v is still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your Tickets here to attend")

}
func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range Bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email address: ")
	fmt.Scan(&emailAddress)

	fmt.Print("Enter the number of tickets : ")
	fmt.Scan(&userTickets)
	return firstName, lastName, emailAddress, userTickets
}
func bookTicket(userTickets int, firstName, lastName, emailAddress string) {
	remainingTickets = remainingTickets - userTickets

	var userData = userData{
		firstName:    firstName,
		lastName:     lastName,
		email:        emailAddress,
		numOfTickets: userTickets,
	}

	Bookings = append(Bookings, userData)
	fmt.Printf("list of booking is %v\n", Bookings)

	fmt.Printf("Thank you %v %v for purchasing %v tickets, a confirmation message will be sent to your mail at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(firstName, lastName, emailAddress string, userTickets int) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, emailAddress)
	fmt.Println("##########")
	wg.Done()
}
