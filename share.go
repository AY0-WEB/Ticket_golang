package main

import "strings"

func validDoc(firstName, lastName, emailAddress string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(emailAddress, "@")
	isValidTicketNum := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNum
}
