package main

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

func main() {
    protocol := "udp"
    buf := make([]byte, 1024)

    fmt.Println("Starting Server - INTERCEPTOR")
    fmt.Println("Waiting Client Choice")

    // Listen to Port
    addr1 := net.UDPAddr{IP:net.ParseIP("127.0.0.1"),Port:9067}
    // Listen to specific Socket - UDP
    conn, _ := net.ListenUDP(protocol, &addr1)
    // Dial to Server 2
    conn_brain, _ := net.Dial(protocol, "127.0.0.1:9068")

    // Infinity Loop
    for {
        // Listen for message to process and write to buffer
        _, remoteaddr, _ := conn.ReadFromUDP(buf)
        user_choice := Choice(string(buf))
        valid_input := ValidateUserChoice(user_choice)
        
        if valid_input{
            // Request choice to BRAIN
            fmt.Fprintf(conn_brain, "Waiting Brain Choice" + "\n")
            
            // Receiving BRAIN choice
            brain_message, _ := bufio.NewReader(conn_brain).ReadString('\n')
            brain_choice := Choice(strings.TrimSuffix(brain_message, "\n"))
            client_response := "Result"

            // Calculate Winner
            if strings.EqualFold(user_choice, brain_choice){
                client_response = DrawMessageBuilder(user_choice, brain_choice)

            }else if strings.EqualFold(user_choice, "Rock"){
                if strings.EqualFold(brain_choice, "Scissor"){
                    client_response = WinMessageBuilder(user_choice, brain_choice)
                }else{
                    client_response = LoseMessageBuilder(user_choice, brain_choice)
                }

            }else if strings.EqualFold(user_choice, "Paper"){
                if strings.EqualFold(brain_choice, "Rock"){
                    client_response = WinMessageBuilder(user_choice, brain_choice)
                }else{
                    client_response = LoseMessageBuilder(user_choice, brain_choice)
                }

            }else{
                if strings.EqualFold(brain_choice, "Paper"){
                    client_response = WinMessageBuilder(user_choice, brain_choice)
                }else{
                    client_response = LoseMessageBuilder(user_choice, brain_choice)
                }

            }
            conn.WriteToUDP([]byte(client_response + "\n"), remoteaddr)
        }else{
            conn.WriteToUDP([]byte("Sorry, Don't know this input, try: R, P or S" + "\n"), remoteaddr)
        }
    }
}

func Choice(input string) string{
	if strings.Contains(strings.ToLower(input),"r"){
		return "Rock"
	}else if strings.Contains(strings.ToLower(input),"p"){
		return "Paper"
    }else if strings.Contains(strings.ToLower(input),"s"){
		return "Scissor"
	}else{
		return "UNKNOWN"
	}
}

func ValidateUserChoice(choice string) bool{
	if strings.EqualFold(choice, "Rock") || strings.EqualFold(choice, "Scissor") || strings.EqualFold(choice, "Paper") {
		fmt.Print("Valid Choice from User\n")
		return true
	}else{
		fmt.Print("ValueError - Unkown Input\n")
		return false
	}
}

func DrawMessageBuilder(user_choice, brain_choice string) string{
	fmt.Println("It's a Draw!" + " <" + user_choice + " vs " + brain_choice + ">")
	return "It's a Draw!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}

func LoseMessageBuilder(user_choice, brain_choice string) string{
	fmt.Println("User Loses!" + " <" + user_choice + " vs " + brain_choice + ">")
	return "You Lose!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}

func WinMessageBuilder(user_choice, brain_choice string) string{
	fmt.Println("User Win!" + " <" + user_choice + " vs " + brain_choice + ">")
	return "You Win!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}
