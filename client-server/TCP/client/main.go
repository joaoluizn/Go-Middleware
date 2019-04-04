package main

import(
	
	"log"
	"net/rpc"
	"fmt"
)

type Item struct{
	First_round string
	Second_round string
	Third_round string
}

func main(){
	var reply Item

	client, err := rpc.DialHTTP("tcp","localhost:4040")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	player_moves := Item{"Scissors", "Rock", "Rock"}

	client.Call("API.Jokenpo", player_moves, &reply)

}