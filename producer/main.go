package main

import (
	"github.com/macduyhai/go-rabbitMQ/logger"
	"github.com/macduyhai/go-rabbitMQ/producer/config"
	"github.com/macduyhai/go-rabbitMQ/producer/router"
)

func main() {
	logger.LogInfor("Server Producer starting ...")
	config := config.NewConfig()
	if config == nil {
		logger.LogError("read config error")
		return
	}
	router := router.NewRouter(config)
	app, err := router.InitGin()
	if err != nil {
		logger.LogError(err.Error())
		return
	}
	err = app.Run(config.PortEngine)
	if err != nil {
		logger.LogError("Run server error:" + err.Error())
	}
}
