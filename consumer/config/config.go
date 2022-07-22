package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/macduyhai/go-rabbitMQ/logger"
)

type Config struct {
	RMQURL       string `env:"RMQ_URL,required"`
	UserQueue    string `env:"USER_QUEUE,required"`
	ProductQueue string `env:"PRODUCT_QUEUE,required"`
}

func NewConfig(files ...string) *Config {
	cfg := Config{}
	err := godotenv.Load(files...)
	if err != nil {
		logger.LogInfor("Env File could not found ")

		panic(err.Error())

	}
	err = env.Parse(&cfg)

	if err != nil {
		logger.LogInfor("Parse config file error:" + err.Error())
		return nil
	}
	return &cfg

}
