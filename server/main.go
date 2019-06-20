package server

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/models"
)

func HandleMQMessage(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s from [%s] queue", d.Body, d.RoutingKey)

		task := models.JudgementTask{}
		result := models.JudgementResult{}

		err := json.Unmarshal(d.Body, &task)
		if err != nil {
			result.Succeed = false
			result.Result = models.SE
			result.ErrorInfo = "The form of task is illegal"
		} else {
			doHandle(task, &result)
			result.Succeed = true
		}
		PublishResult(result)
	}
}

func doHandle(task models.JudgementTask, result *models.JudgementResult) {
	result.Id = task.Id

	status, errorInfo := judger.Run(task)
	result.Result = status
	result.ErrorInfo = errorInfo
}

func PublishResult(result models.JudgementResult) {
	err := ch.Publish(
		"",
		"go-docker-judger-callback",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        result.ToJsonString(),
		})

	if err != nil {
		log.Println("Fail to handle middle => " + result.Id)
	}
}
