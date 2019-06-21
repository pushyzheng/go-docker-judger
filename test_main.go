package main

import (
	"log"
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/models"
)

func main() {
	//reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(os.Stdout, reader)
	//
	//GetAllImage(cli, ctx)
	//GetAllContainer(cli, ctx)
	//
	//CreateHelloWorld(cli, ctx)

	conf.InitConfig()

	task := models.JudgementTask{}
	task.UserId = "123"
	task.ProblemId = 1
	task.TimeLimit = 1
	task.MemoryLimit = 30

	result, err := judger.VerifyAnswer(task)
	if err != nil{
		panic(err)
	}

	log.Println("status: ", result.Status)
	log.Println("Last input: ", result.LastInput)
	log.Println("Last output: ", result.LastOutput)
	log.Println("Expected output: ", result.ExpectedOutput)

}