package conf

import (
	"github.com/go-ini/ini"
	"log"
)

var Container = &DockerContainerConfig{}
var Volume = &VolumePathConfig{}
var RabbitMQ = &RabbitMQConfig{}

var cfg *ini.File

func InitConfig() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	mapTo("volume", Volume)
	mapTo("container", Container)
	mapTo("rabbitmq", RabbitMQ)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}