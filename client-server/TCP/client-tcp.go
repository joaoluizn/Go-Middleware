package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// main Client main function - Jokenpo TCP
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9067")
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
