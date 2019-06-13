package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

// main Server Two main function responsible for process random choice.
func main() {
	fmt.Println("Starting Server - BRAIN")
	choices := make([]string, 0)
	choices = append(choices,
		"r",
		"p",
		"s",
	)
	// Generate Random seed
	rand.Seed(time.Now().Unix())

	// Listen to Port
	ln, err := net.Listen("tcp", ":9068")
	checkError(err)

	// Accepting connection
	conn, err := ln.Accept()
	checkError(err)
	defer conn.Close()

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)
		fmt.Print("Message Received: ", string(message))

		newMessage := string(choices[rand.Intn(len(choices))])
		fmt.Print("Random Brain Choice: ", newMessage+"\n")

		conn.Write([]byte(newMessage + "\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
