package main

import (
	"fmt"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        fmt.Printf("%s: %s", msg, err)
    }
}

func main() {
	fmt.Println("Starting Server - BRAIN")
	choices := make([]string, 0)
	choices = append(choices,
		"r",
		"p",
		"s",
		"p",
		"r",
	)
	
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    chSend, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chSend.Close()

    qSend, err := chSend.QueueDeclare(
        "server-2", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")


    chReceive, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chReceive.Close()

    qReceive, err := chReceive.QueueDeclare(
        "server-1", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

  	i := 0
  	for {
    	// Wait Reply from Server 1
        msgs, err := chReceive.Consume(
            qReceive.Name, // queue
            "",     // consumer
            true,   // auto-ack
            false,  // exclusive
            false,  // no-local
            false,  // no-wait
            nil,    // args
        )
        failOnError(err, "Failed to register a consumer")

        for range msgs{
            fmt.Printf("Received")
        }

		// Random choice
    	newmessage := string(choices[i%5])

        err = chSend.Publish(
            "",     // exchange
            qSend.Name, // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(newmessage),
        })
        failOnError(err, "Failed to publish a message")
  	}
}