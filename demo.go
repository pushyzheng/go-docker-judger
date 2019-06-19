package main

import (
	"pushy.site/go-docker-judger/conf"
	"pushy.site/go-docker-judger/judger"
)

func main() {
	judger.InitCore()
	conf.InitConfig()

	judger.Run()
}