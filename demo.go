package main

import (
	"log"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/models"
)

func Run() {
	conf.InitConfig()
	judger.InitCore()

	task := models.JudgementTask{}
	task.UserId = "123"
	task.ProblemId = 1
	task.TimeLimit = 1
	task.MemoryLimit = 30
	task.Language = "java"

	result := models.JudgementResult{}
	judger.Run(task, &result)

	log.Println("Status: ", result.Status)
	log.Println("errorInfo: ", result.ErrorInfo)

	log.Println("Last input: ", result.LastInput)
	log.Println("Last output: ", result.LastOutput)
	log.Println("Expected output: ", result.ExpectedOutput)
}

func main() {
	Run()
}