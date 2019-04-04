package main

import(
	"net"
	"net/rpc"
	"net/http"
	"log"
	"fmt"
)

type Item struct{
	First_round string
	Second_round string
	Third_round string
}

type API int

var count = 0;

func (a *API) Jokenpo(item Item, reply *string) error{
	//database = append(database, item)
	computer_moves := Item{"Rock", "Scissor", "Paper"}
	
	//PRIMEIRO ROUND
	fmt.Println("First round:\nYou: " + computer_moves.First_round + "\nComputer: Rock")
	if computer_moves.First_round=="Rock"{
		fmt.Println("Tie!")
	}else if computer_moves.First_round=="Scissor"{
		fmt.Println("You Lose!")
		count = count - 1
	}else{
		fmt.Println("You Win!")
		count = count + 1
	}

	//SEGUNDO ROUND
	fmt.Println("Second round:\nYou: " + computer_moves.First_round + "\nComputer: Scissors")
	if computer_moves.Second_round=="Rock"{
		fmt.Println("You Win !")
		count = count + 1
	}else if computer_moves.Second_round=="Scissor"{
		fmt.Println("Tie!")
	}else{
		fmt.Println("You Lose!")
		count = count - 1
	}

	//TERCEIRO ROUND
	fmt.Println("Third round:\nYou: " + computer_moves.First_round + "\nComputer: Scissors")
	if computer_moves.Third_round=="Rock"{
		fmt.Println("You Lose!")
		count = count - 1
	}else if computer_moves.Third_round=="Scissor"{
		fmt.Println("You Win!")
		count = count + 1
	}else{
		fmt.Println("Tie!")
	}

	var retorno string

	if count > 0 {
		retorno = "Voce venceu!"
	}else if count == 0{
		retorno = "voce perdeu"
	}else{
		retorno = "EMPATE"
	}

	*reply = retorno
	return nil
}

func main(){

	var api = new(API)
	err := rpc.Register(api)

	if err != nil{
		log.Fatal("error registering API", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil{
		log.Fatal("errr!")
	}

	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil{ 
		log.Fatal("error serving ", err)
	}

}

