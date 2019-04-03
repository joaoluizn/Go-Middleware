package main

import (
    "net"
    "fmt"
    "bufio"
    "os"
    "time"
)

func main() {
    var i int

    fmt.Printf("How many times it will be executed?")
    _, err := fmt.Scanf("%d", &i)

    fmt.Printf("%d\n", i)

    choices := make([]string, 0)
    choices = append(choices,
        "r",
        "p",
        "s",
    )

    f, err := os.Create(fmt.Sprintf("%dtcp.txt",i))
    if err != nil {
        fmt.Println(err)
        return
    }

    // Dial to specified Socket
    conn, _ := net.Dial("tcp", "127.0.0.1:9067")

    for j :=0; j<i; j++{
        // Receive input
        text := string(choices[j%3])

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