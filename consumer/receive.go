package main

import (
	"log"

	"github.com/leepuppychow/rabbitMQ-practice/utils"
)

func main() {
	conn := utils.ConnectToRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"testChan1", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to register consumer")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Waiting for a message. To exit press CTRL+C")
	<-forever	// Note: here waiting to receive from an empty channel, so main goroutine is blocked forever
}
