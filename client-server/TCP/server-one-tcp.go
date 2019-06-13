package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// main Server one main function responsible to judge choices from client and server two
func main() {
	protocol := "tcp"

	fmt.Println("Starting Server - INTERCEPTOR")
	fmt.Println("Waiting Client Choice")

	ln, err := net.Listen(protocol, ":9067")
	checkError(err)

	// Accepting connection
	conn, err := ln.Accept()
	checkError(err)

	// Dial to Server 2
	connBrain, err := net.Dial(protocol, "127.0.0.1:9068")
	checkError(err)

	defer conn.Close()
	defer connBrain.Close()

	for {
		// Listen for message to process with newline
		message, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)

		userChoice := choice(strings.TrimSuffix(message, "\n"))
		isInputValid := validateUserChoice(userChoice)

		if isInputValid {
			// Request choice to BRAIN
			fmt.Fprintf(connBrain, "Waiting Brain Choice"+"\n")

			// Receiving BRAIN choice
			brainMessage, err := bufio.NewReader(connBrain).ReadString('\n')
			checkError(err)

			brainChoice := choice(strings.TrimSuffix(brainMessage, "\n"))
			clientResponse := "Result"

			// Calculate Winner
			if strings.EqualFold(userChoice, brainChoice) {
				clientResponse = drawMessageBuilder(userChoice, brainChoice)

			} else if strings.EqualFold(userChoice, "Rock") {
				if strings.EqualFold(brainChoice, "Scissor") {
					clientResponse = winMessageBuilder(userChoice, brainChoice)
				} else {
					clientResponse = loseMessageBuilder(userChoice, brainChoice)
				}

			} else if strings.EqualFold(userChoice, "Paper") {
				if strings.EqualFold(brainChoice, "Rock") {
					clientResponse = winMessageBuilder(userChoice, brainChoice)
				} else {
					clientResponse = loseMessageBuilder(userChoice, brainChoice)
				}

			} else {
				if strings.EqualFold(brainChoice, "Paper") {
					clientResponse = winMessageBuilder(userChoice, brainChoice)
				} else {
					clientResponse = loseMessageBuilder(userChoice, brainChoice)
				}

			}
			conn.Write([]byte(clientResponse + "\n"))
		} else {
			conn.Write([]byte("Sorry, Don't know this input, try: R, P or S" + "\n"))
		}
	}
}

// choice Parse Jokenpo choices from letters
func choice(a string) string {
	if strings.EqualFold(a, "r") {
		return "Rock"
	} else if strings.EqualFold(a, "p") {
		return "Paper"
	} else if strings.EqualFold(a, "s") {
		return "Scissor"
	} else {
		return "UNKNOWN"
	}
}

// validateUserChoice Validates Jokenpo choices from String
func validateUserChoice(choice string) bool {
	if strings.EqualFold(choice, "Rock") || strings.EqualFold(choice, "Scissor") || strings.EqualFold(choice, "Paper") {
		fmt.Print("Valid Choice from User\n")
		return true
	}
	fmt.Print("ValueError - Unkown Input\n")
	return false
}

// drawMessageBuilder Build default Draw message
func drawMessageBuilder(userChoice, brainChoice string) string {
	fmt.Println("It's a Draw!" + " <" + userChoice + " vs " + brainChoice + ">")
	return "It's a Draw!" + " User > " + userChoice + " vs " + brainChoice + " < Brain"
}

// loseMessageBuilder Build default lose message
func loseMessageBuilder(userChoice, brainChoice string) string {
	fmt.Println("User Loses!" + " <" + userChoice + " vs " + brainChoice + ">")
	return "You Lose!" + " User > " + userChoice + " vs " + brainChoice + " < Brain"
}

// winMessageBuilder Build default win message
func winMessageBuilder(userChoice, brainChoice string) string {
	fmt.Println("User Win!" + " <" + userChoice + " vs " + brainChoice + ">")
	return "You Win!" + " User > " + userChoice + " vs " + brainChoice + " < Brain"
}

// checkError Simple Error Handler
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
