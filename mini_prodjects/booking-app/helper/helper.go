package helper

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidName, isValidTicketNumber
}

func HandlePersonifiedErrors(isValidEmail bool, isValidName bool, isValidTicketNumber bool) {
	if !isValidEmail {
		fmt.Println("Your email is invalid and does not contain '@'")
	}

	if !isValidName {
		fmt.Println("Your name is too short")
	}

	if !isValidTicketNumber {
		fmt.Println("Your ticket number is invalid")
	}
}
