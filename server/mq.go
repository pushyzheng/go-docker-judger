package server

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"pushy.site/go-docker-judger/conf"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQ() {
	var err error

	log.Println("Rabbitmq connect URL: " + conf.RabbitMQ.GetURL())

	conn, err = amqp.Dial(conf.RabbitMQ.GetURL())
	if err != nil {
		panic(err)
	}

	ch, err = conn.Channel()
	if err != nil {
		panic(err)
	}
}

func StartConsuming() {
	queue, err := ch.QueueDeclare(
		"go-docker-judger", false, false,
		false, false, nil,
	)
	if err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("Declare queue[%s] successfully", queue.Name))

	msgs, err := ch.Consume(
		"go-docker-judger",
		"", true, false,
		false, false, nil,
	)
	if err != nil {
		panic(err)
	}
	forever := make(chan bool)

	go HandleMQMessage(msgs)

	log.Printf(" [*] Start consuming judger message queue. To exit press CTRL+C")
	<-forever
}

func CloseMQConnection() {
	conn.Close()
	ch.Close()
}