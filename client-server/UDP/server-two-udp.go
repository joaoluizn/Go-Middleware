package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

// main Server Two main function responsible for process random choice.
func main() {
	buf := make([]byte, 1024)
	addr2 := net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 9068}

	fmt.Println("Starting Server - BRAIN")
	choices := make([]string, 0)
	choices = append(choices,
		"r",
		"p",
		"s",
	)
	// Generate Random seed
	rand.Seed(time.Now().Unix())
	conn, err := net.ListenUDP("udp", &addr2)

	checkError(err)
	defer conn.Close()

	for {
		_, remoteaddr, err := conn.ReadFromUDP(buf)
		checkError(err)

		fmt.Print("Message Received: ", string(buf))

		brainChoice := string(choices[rand.Intn(len(choices))])
		fmt.Print("Random Brain Choice: ", brainChoice+"\n")

		conn.WriteToUDP([]byte(brainChoice+"\n"), remoteaddr)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
