package main

import (
	"log"
	"os"

	"github.com/leepuppychow/rabbitMQ-practice/utils"
	"github.com/streadway/amqp"
)

func main() {
	conn := utils.ConnectToRabbitMQ()
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"queue1", 	 // name
		true,        // durable (in case the RabbitMQ server crashes)
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	utils.FailOnError(err, "Failed to declare a queue")

	body := utils.BodyFrom(os.Args)
	err = ch.Publish(
		"",		// Note: here we are just using a default exchange
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,	// tries to save message on disk in case it is lost (not perfect persistence though)
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	utils.FailOnError(err, "Failed to publish message")
	log.Printf("Sent %s", body)
}
