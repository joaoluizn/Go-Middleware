package main

import (
	"net"
    "net/rpc"
    "net/http"
    "log"
)

type Item struct{
    First_round string
    Second_round string
    Third_round string
}

type API int

func(a *API) PCMoves(item Item, reply *Item) error{
    computer_moves  := Item{"Rock", "Scissor", "Paper"}
	*reply = computer_moves 
	return nil
}

func main() {
	
	var api = new(API)
    err := rpc.Register(api)

    if err != nil{
        log.Fatal("error registering API", err)
    }

    rpc.HandleHTTP()

    listener, err := net.Listen("tcp", ":1313")

    if err != nil{
        log.Fatal("errr!")
    }

    log.Printf("serving rpc on port %d", 1313)
    err = http.Serve(listener, nil)
    if err != nil{ 
        log.Fatal("error serving ", err)
    }

	/*fmt.Println("Starting Server02 - Computer player")
	choices := make([]string, 0)
	choices = append(choices,
		"r",
		"p",
		"s",
		"p",
		"r",
	)*/
	

	/*// Generate Random seed
	rand.Seed(time.Now().Unix())
	// Listen to Port
  	ln, _ := net.Listen("tcp", ":9068")
  	// Accepting connection
  	conn, _ := ln.Accept()
  	// Infinity Loop
  	i := 0*/
  	
  	/*for {
    	message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received: ", string(message))
		// Random choice
    	newmessage := string(choices[i%5])
    	conn.Write([]byte(newmessage + "\n"))
    	i++;
  	}*/
}