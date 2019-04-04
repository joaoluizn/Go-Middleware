package main

import(
	"fmt"
	"log"
	"net/rpc"
	
)

type Item struct{
	First_round string
	Second_round string
	Third_round string
}

func main(){
	var reply string

	client, err := rpc.DialHTTP("tcp","localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	player_moves := Item{"Scissors", "Rock", "Rock"}

	client.Call("API.Jokenpo", player_moves, &reply)

	fmt.Println(reply)

}