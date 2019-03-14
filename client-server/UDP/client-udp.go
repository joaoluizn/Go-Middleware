package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "time"
    "math/rand"
)

func main() {

    var i int
    _, err := fmt.Scanf("%d", &i)

    fmt.Printf("%d\n", i)

    choices := make([]string, 0)
    choices = append(choices,
        "r",
        "p",
        "s",
    )

    f, err := os.Create(fmt.Sprintf("%dudp.txt",i))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Dial to specified Socket
    ServerAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9067")
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	CheckError(err)

    conn, _ := net.DialUDP("udp", LocalAddr, ServerAddr)
    
    for j :=0; j<i; j++{
        text := string(choices[rand.Intn(len(choices))])

        start := time.Now()

        // Send response to socket
        fmt.Fprintf(conn, text + "\n")

        // Wait Reply from Server 1
        message, _ := bufio.NewReader(conn).ReadString('\n')
        fmt.Println("Result: " + message)

        elapsed := time.Since(start)

        f.WriteString(fmt.Sprintf("%f\n",elapsed))
    }
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}