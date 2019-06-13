package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// main Client main function - Jokenpo
func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9067")
	defer conn.Close()
	for {
		// Receive input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("(Choose R, P or S): ")
		text, _ := reader.ReadString('\n')

		// Send response to socket
		fmt.Fprintf(conn, text+"\n")

		// Wait Reply from Server 1
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Result: " + message)
	}
}
