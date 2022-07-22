package main

import (
	"github.com/macduyhai/go-rabbitMQ/consumer/config"
	"github.com/macduyhai/go-rabbitMQ/consumer/service"
	"github.com/macduyhai/go-rabbitMQ/logger"
)

func main() {
	logger.LogInfor("Server consumer is starting ...")
	config := config.NewConfig()
	if config == nil {
		logger.LogError("read config error")
		return
	}
	rbmq := service.InitConnectionRB(config.RMQURL, config.UserQueue)
	rbservice := service.GetRMQConsumerService(rbmq)
	err := rbservice.HanderConsumer()
	if err != nil {
		logger.LogError(err)
		return
	}
}
