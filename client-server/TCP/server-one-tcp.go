package main

import (
    "fmt"
    "strings"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        fmt.Printf("%s: %s", msg, err)
    }
}

var userWins int = 0
var compWins int =0

func main() {

    fmt.Println("Starting Server - INTERCEPTOR")
    fmt.Println("Waiting Client Choice")

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    chSend, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chSend.Close()

    qSend, err := chSend.QueueDeclare(
        "server-1", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")


    chReceive1, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chReceive1.Close()

    qReceive1, err := chReceive1.QueueDeclare(
        "client", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    chReceive2, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chReceive2.Close()

    qReceive2, err := chReceive2.QueueDeclare(
        "server-2", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")    


    // Infinity Loop
    for {
        message := "";

        // Listen for message to process with newline
         msgs, err := chReceive1.Consume(
            qReceive1.Name, // queue
            "",     // consumer
            true,   // auto-ack
            false,  // exclusive
            false,  // no-local
            false,  // no-wait
            nil,    // args
        )
        failOnError(err, "Failed to register a consumer")
            
        for d := range msgs{
            message = string(d.Body)
        }

        user_choice := Choice(strings.TrimSuffix(message, "\n"))
        valid_input := ValidateUserChoice(user_choice)
        
        if valid_input{
            // Request choice to BRAIN
            fmt.Printf("Waiting Brain Choice" + "\n")
            
            // Receiving BRAIN choice
            msgs, err := chReceive2.Consume(
                qReceive2.Name, // queue
                "",     // consumer
                true,   // auto-ack
                false,  // exclusive
                false,  // no-local
                false,  // no-wait
                nil,    // args
            )
            failOnError(err, "Failed to register a consumer")

            brain_choice := ""

            for d := range msgs{
                brain_choice = string(d.Body)
            }

            client_response := ""
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
            err = chSend.Publish(
                "",     // exchange
                qSend.Name, // routing key
                false,  // mandatory
                false,  // immediate
                amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte(client_response),
            })
            failOnError(err, "Failed to publish a message")

        }else{
            err = chSend.Publish(
                "",     // exchange
                qSend.Name, // routing key
                false,  // mandatory
                false,  // immediate
                amqp.Publishing{
                ContentType: "text/plain",
                Body:        []byte("Sorry, Don't know this input, try: R, P or S" + "\n"),
            })
            failOnError(err, "Failed to publish a message")
        }
    }
}

func Choice(a string) string{
	if strings.EqualFold(a, "r"){
		return "Rock"
	}else if strings.EqualFold(a, "p"){
		return "Paper"
	}else if strings.EqualFold(a, "s"){
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
    userWins++
    compWins++
	return "It's a Draw!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}

func LoseMessageBuilder(user_choice, brain_choice string) string{
	fmt.Println("User Loses!" + " <" + user_choice + " vs " + brain_choice + ">")
    compWins++
	return "You Lose!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}

func WinMessageBuilder(user_choice, brain_choice string) string{
	fmt.Println("User Win!" + " <" + user_choice + " vs " + brain_choice + ">")
    userWins++
	return "You Win!" + " User > " + user_choice + " vs " + brain_choice + " < Brain"
}
