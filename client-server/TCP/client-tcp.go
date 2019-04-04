package main

import (
    "fmt"
    "os"
    "time"
    "github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
    if err != nil {
        fmt.Printf("%s: %s", msg, err)
    }
}

func main() {
    var i int

    fmt.Printf("How many times it will be executed?")
    _, err := fmt.Scanf("%d", &i)

    fmt.Printf("%d\n", i)

    choices := make([]string, 0)
    choices = append(choices,
        "r",
        "p",
        "s",
    )

    f, err := os.Create(fmt.Sprintf("%dtcp.txt",i))
    if err != nil {
        fmt.Println(err)
        return
    }

    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    chSend, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer chSend.Close()

    qSend, err := chSend.QueueDeclare(
        "client", // name
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
        "server-1-client", // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    failOnError(err, "Failed to declare a queue")

    for j :=0; j<i; j++{
        // Receive input

        start := time.Now()

        body := string(choices[j%3])

        err = chSend.Publish(
            "",     // exchange
            qSend.Name, // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing{
            ContentType: "text/plain",
            Body:        []byte(body),
        })
        failOnError(err, "Failed to publish a message")

        fmt.Printf("Sent\n")

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

        for d := range msgs{
            fmt.Printf("Result: %s", d.Body)
        }

        elapsed := time.Since(start)

        f.WriteString(fmt.Sprintf("%f\n",elapsed))

    }
}