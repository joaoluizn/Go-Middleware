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
		"p",
		"r",
	)
	// Generate Random seed
	rand.Seed(time.Now().Unix())
	// Listen to Port
  	ln, _ := net.Listen("tcp", ":9068")
  	// Accepting connection
  	conn, _ := ln.Accept()
  	// Infinity Loop
  	i := 0
  	for {
    	message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		// Random choice
    	newmessage := string(choices[i%5])
    	conn.Write([]byte(newmessage + "\n"))
    	i++;
  	}
}