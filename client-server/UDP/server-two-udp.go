package main

import (
	"net"
	"fmt"
	"math/rand"
	"time"
)

func main() {
    buf := make([]byte, 1024)
	addr2 := net.UDPAddr{IP:net.ParseIP("127.0.0.1"),Port:9068}
	fmt.Println("Starting Server - BRAIN")
	// Listen to Port
	choices := make([]string, 0)
	choices = append(choices,
		"r",
		"p",
		"s",
	)
	// Generate Random seed
	rand.Seed(time.Now().Unix())
	conn, _ := net.ListenUDP("udp", &addr2)

	// Infinity Loop
  	for {
		_, remoteaddr, _ := conn.ReadFromUDP(buf)
		fmt.Print("Message Received: ", string(buf))
		// Random choice
		brain_choice := string(choices[rand.Intn(len(choices))])
		fmt.Print("Choice: ", brain_choice + "\n")

		conn.WriteToUDP([]byte(brain_choice + "\n"), remoteaddr)
  	}
}