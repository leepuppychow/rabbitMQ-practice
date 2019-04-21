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

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	body := utils.BodyFrom(os.Args)
	err = ch.Publish(
		"logs",		// exchange
		"",				// routing key (ignored with fanout type exchange)
		false,		// mandatory
		false,		// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	utils.FailOnError(err, "Failed to publish message")
	log.Printf("Sent %s", body)
}
