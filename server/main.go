package server

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"go-docker-judger/judger"
	"go-docker-judger/models"
	"log"
)

// 处理服务端通过消息队列发送的判题任务
func HandleMQMessage(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		log.Printf("Received a message: %s from [%s] queue", d.Body, d.RoutingKey)

		task := models.JudgementTask{}
		result := models.JudgementResult{}

		err := json.Unmarshal(d.Body, &task)
		if err != nil {
			result.Succeed = false
			result.Status = models.SE
			result.ErrorInfo = "The form of task is illegal"
		} else {
			doHandle(task, &result)
			result.Succeed = true
		}
		PublishResult(result)
	}
}

// 真正处理消息的逻辑方法
func doHandle(task models.JudgementTask, result *models.JudgementResult) {
	result.Id = task.Id
	judger.Run(task, result) // 判题核心
}

// 处理完判题任务之后，发布到消息队列
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
