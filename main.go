package main

import (
	"go-docker-judger/conf"
	"go-docker-judger/judger"
	"go-docker-judger/server"
)

func main() {
	conf.InitConfig()
	judger.InitCore()

	// 初始化 RabbitMQ 相关
	server.InitRabbitMQ()
	server.StartConsuming()
	defer server.CloseMQConnection()
}
