package main

import (
	"net"
	"fmt"
	"bufio"
	"math/rand"
	"time"
)

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
  	ln, _ := net.Listen("tcp", ":9068")
  	// Accepting connection
  	conn, _ := ln.Accept()
  	// Infinity Loop
  	for {
    	message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		// Random choice
    	newmessage := string(choices[rand.Intn(len(choices))])
    	conn.Write([]byte(newmessage + "\n"))
  	}
}