package server

import (
	"github.com/streadway/amqp"
	"log"
)

func HandleMQMessage(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s from [%s] queue", d.Body, d.RoutingKey)
	}
}