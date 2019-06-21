package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/models"
	"pushy.site/go-docker-judger/utils"
)

func Run() {
	conf.InitConfig()
	judger.InitCore()

	task := models.JudgementTask{}
	task.UserId = "123"
	task.ProblemId = 1
	task.TimeLimit = 1
	task.MemoryLimit = 30

	//result, _ := judger.Run(task)
	//fmt.Println("Status: " + result.Status)
}

func main() {
	//Run()

	fileBytes, err := ioutil.ReadFile("e:/usr/answers/answer_1.txt")
	if err != nil {
		panic(err)
	}

	var b bytes.Buffer
	b.Write(fileBytes)

	line := utils.GetLineByBytes(b.Bytes(), 3)
	fmt.Println(line)

	count := utils.GetLineCountByBytes(b.Bytes())
	fmt.Println("count:", count)
}