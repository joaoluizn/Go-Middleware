package main

import(
    "fmt"
    "log"
    "net/rpc"
    "time"
)

type Item struct{
    First_round string
    Second_round string
    Third_round string
}

func main(){
    
    start := time.Now()

    var reply string

    client, err := rpc.DialHTTP("tcp","localhost:4040")
    if err != nil {
        log.Fatal("Connection error: ", err)
    }

    player_moves := Item{"Scissors", "Rock", "Rock"}

    count := 0

    for count <= 10000 {
        client.Call("API.Jokenpo", player_moves, &reply)
        count = count+1
    }

    fmt.Println(reply)

    duration := time.Since(start)
    fmt.Println("\nDesempenho: ", duration)
}