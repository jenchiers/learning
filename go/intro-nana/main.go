package main

import (
	"booking-app/helper"
	"fmt"
	"time"
)

const conferenceTickets = 50

// alternative way of assinging a value to a variable when it's created is not possible when defining package level variables
var conferenceName = "Go Conference"

// let Go infer the type of the variable, unless you want to assign a different type than what go would assign
var remainingTickets uint = 50

// slice is an abstraction of an array and has a dynamic size, so mostly a slice is a better option to use instead of array
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

// a struct can hold values of different types
type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
	//newsLetter bool
}

// create a waitgroup to prevent the main application from shutting down before all threads are completed, which it does by default
// this code is added as an example, but we don't need it because there is an endless for loop and the program is ended manually
// var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	/*
		fmt.Printf("conferenceTickets is %T\n", conferenceTickets)
		fmt.Printf("remainingTickets is %T\n", remainingTickets)
		fmt.Printf("conferenceName is %T\n", conferenceName)
	*/

	// print out the actual value of a variable
	fmt.Println(remainingTickets)
	// print out the memory location of the remainingTickets variable, which is a memory value that points to the memory location of the variable (pointer)
	fmt.Println(&remainingTickets)

	// you can't mix types in an array, an array has a fixed size
	// var bookings = [50]string{} or:
	// var bookings [50]string

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// add the number of goroutines to wait for
			// wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("These people have booked tickets: %v.\n", firstNames)

			// alternative, the condition can be added directly behind the if statement, here we create a bool first
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				// end the loop/program
				fmt.Println("Our conference is fully booked.")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Firstname or lastname is invalid: %v %v.\n", firstName, lastName)
			}
			if !isValidEmail {
				fmt.Printf("Email does not contain @ %v.\n", email)
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of tickets bought is invalid %v.\n", userTickets)
			}

			// continue to the next iteration
			// continue
		}
	}
	// this waits for all threads to be done that are added before, a wg.Done() function has to be added to the function that is executed in a goroutine
	// wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v conference tickets and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
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
	// ask for the user's information
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets you want to buy:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// variable types must be the same to use them in calculations (both are uint for example)
	remainingTickets -= userTickets

	// create a map for a user + booking, data types can not be mixed in a map, so an uint for example needs to be converted
	var bookingDetails = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	// bookingDetails["firstName"] = firstName
	// bookingDetails["lastName"] = lastName
	// bookingDetails["email"] = email
	// bookingDetails["userTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, bookingDetails)
	fmt.Printf("List of bookings %v.\n", bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value of the slice: %v\n", bookings[0])
	// fmt.Printf("The type of the slice: %T\n", bookings)
	// fmt.Printf("The size of the slice: %v\n", len(bookings))

	fmt.Printf("%v %v booked %v conference tickets, the tickets will be sent to: %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v remaining tickets of the %v total tickets.\n", remainingTickets, conferenceTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v.\n", userTickets, firstName, lastName)
	fmt.Println("------------------------------------")
	fmt.Printf("Sending ticket:\n %v \nto %v.\n", ticket, email)
	fmt.Println("------------------------------------")
	// wg.Done()
}
