package main

import (
	"fmt"
	"time" 
)

const conferenceTickets int = 50

var conferenceName = "OEW Conference"
var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	greetUser()

	for {
		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isvalidTicketNumber := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isvalidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			go sendTickets(userTickets, firstName, lastName, email)

			firstNames := getprintFirstNames()
			fmt.Printf("the first nanes of our bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("our conference is booked out, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name is incorrect")
			}
			if !isValidEmail {
				fmt.Println("email you entered is incorrect")
			}
			if !isvalidTicketNumber {
				fmt.Println("number of tickets is incorrect")
			}

		}
	}

}

func greetUser() {
	fmt.Printf("welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getprintFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last Name name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets name")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user

	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf(" list of bookings is %v\n", bookings)

	fmt.Printf("thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

func sendTickets(userTickets uint, firstNames string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstNames, lastName)
	fmt.Println("############")
	fmt.Printf("sending tickets:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("############")

}
