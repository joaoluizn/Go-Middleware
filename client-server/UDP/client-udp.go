package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
)

func main() {
    // Dial to specified Socket
    conn, _ := net.Dial("udp", "127.0.0.1:9067")
    for {
        // Receive input
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("(Choose R, P or S): ")
        text, _ := reader.ReadString('\n')

        // Send response to socket
        fmt.Fprintf(conn, text + "\n")

        // Wait Reply from Server 1
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Result: " + message)
    }
}