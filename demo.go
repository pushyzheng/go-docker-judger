package main

import (
	"fmt"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/models"
)

func main() {
	conf.InitConfig()
	judger.InitCore()

	task := models.JudgementTask{}
	task.UserId = "123"
	task.ProblemId = 1
	task.TimeLimit = 1
	task.MemoryLimit = 32

	result, errorInfo := judger.Run(task)
	fmt.Println(result)
	fmt.Println(errorInfo)
}