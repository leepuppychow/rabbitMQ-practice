package main

import (
	"bytes"
	"log"
	"time"

	"github.com/leepuppychow/rabbitMQ-practice/utils"
)

func main() {
	conn := utils.ConnectToRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue1", 	 // name
		true,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	messages, err := ch.Consume(
		q.Name,	// queue
		"",			// consumer
		false,	// sets auto-acknowledgement to false
		false,	// exclusive
		false,	// no-local
		false,	// no-wait
		nil,		// args
	)
	utils.FailOnError(err, "Failed to register consumer")

	forever := make(chan bool)

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)	// Here we manually acknowledge worker is done with task
		}
	}()

	log.Printf("Waiting for a message. To exit press CTRL+C")
	<-forever // Note: here waiting to receive from an empty channel, so main goroutine is blocked forever
}
