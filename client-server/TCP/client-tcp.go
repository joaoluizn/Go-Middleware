package main

import (
	"log"
	"time"
	"fmt"
	"os"
	// "sync"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// Start of testing code
	
	var i int
    log.Printf("How many times it will be executed?")
    _, err := fmt.Scanf("%d", &i)

    log.Printf("%d\n", i)

    choices := make([]string, 0)
    choices = append(choices,
        "r",
        "p",
        "s",
    )

    f, err := os.Create(fmt.Sprintf("%dtcp.txt",i))
    if err != nil {
        log.Fatalf("%s", err)
        return
    }
	// End of testing code

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"server.one", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	resp, err := ch.QueueDeclare(
		"response-s1", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")


	msgs, err := ch.Consume(
		resp.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	start := time.Now()

    for j :=0; j<i; j++{
		start = time.Now()
		body := string(choices[j%3])

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
				MessageId: "client",
			})

		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")
	}
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
		for d := range msgs{
			log.Printf(">> Received: %s", string(d.Body))
			elapsed := time.Since(start)
			f.WriteString(fmt.Sprintf("%f\n",elapsed))	
		}
	// }()

	// wg.Wait()
	// elapsed := time.Since(start)
	// f.WriteString(fmt.Sprintf("%f\n",elapsed))
}