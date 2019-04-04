package main

import (
    "fmt"
    "strings"
    "log"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        log.Printf("%s: %s", msg, err)
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


    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    receive_one_q, err := ch.QueueDeclare(
        "server.one", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    response_one_q, err := ch.QueueDeclare(
        "response-s1", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    send_two_q, err := ch.QueueDeclare(
        "server.two", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    // Listen from client
    msgs_one, err := ch.Consume(
        receive_one_q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    failOnError(err, "Failed to register a consumer")


    // Infinity Loop
    forever := make(chan bool)

    go func() {
        message := "";
        var user_choice string = ""
        var brain_choice string = ""

        for d := range msgs_one{
            message = string(d.Body)
            messageType := string(d.MessageId)
            
            log.Printf("Messagem: %s & type: %s", message, messageType)
            if strings.Contains(messageType, "client"){
                log.Printf("Log de Client")
                user_choice = Choice(strings.TrimSuffix(message, "\n"))
                valid_input := ValidateUserChoice(user_choice)

                if valid_input{
                    user_choice = Choice(strings.TrimSuffix(message, "\n"))

                    err = ch.Publish(
                        "",     // exchange
                        send_two_q.Name, // routing key
                        false,  // mandatory
                        false,  // immediate
                        amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte("Your Turn"),
                    })
                }else{
                    err = ch.Publish(
                        "",     // exchange
                        response_one_q.Name, // routing key
                        false,  // mandatory
                        false,  // immediate
                        amqp.Publishing{
                        ContentType: "text/plain",
                        Body:        []byte("Sorry, Don't know this input, try: R, P or S" + "\n"),
                    })
                    failOnError(err, "Failed to publish a message")
                }
            }else{
                // Request choice to BRAIN
                log.Printf("Waiting Brain Choice" + "\n")
        
                log.Printf("brain_chose: %s", message)
                brain_choice = message
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
                err = ch.Publish(
                    "",     // exchange
                    response_one_q.Name, // routing key
                    false,  // mandatory
                    false,  // immediate
                    amqp.Publishing{
                    ContentType: "text/plain",
                    Body:        []byte(client_response),
                })
                failOnError(err, "Failed to publish a message")
                
            }

        }

    }()
	<-forever

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
