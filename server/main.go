package server

import (
	"github.com/streadway/amqp"
	"log"
)

func HandleMQMessage(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s from [%s] queue", d.Body, d.RoutingKey)

		body := "WA"

		err := ch.Publish(
			"",
			"go-docker-judger-callback",
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Println("Fail to handle task")
		}
	}
}
