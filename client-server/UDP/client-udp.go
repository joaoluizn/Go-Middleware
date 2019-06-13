package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// main Client main function - Jokenpo UDP
func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9067")
	checkError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	checkError(err)

	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	checkError(err)

	defer conn.Close()

	for {
		// Receive input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("(Choose R, P or S): ")
		text, err := reader.ReadString('\n')
		checkError(err)

		// Send response to socket
		fmt.Fprintf(conn, text+"\n")

		// Wait Reply from Server 1
		message, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)

		fmt.Println("Result: " + message)
	}
}

// checkError Simple Error Handler
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
