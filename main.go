package main

import (
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
	"pushy.site/go-docker-judger/server"
)


func main() {
	conf.InitConfig()
	judger.InitCore()

	// 初始化 RabbitMQ 相关
	server.InitRabbitMQ()
	server.StartConsuming()
	defer server.CloseMQConnection()
}
