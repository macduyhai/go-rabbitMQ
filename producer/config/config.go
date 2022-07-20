package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/macduyhai/go-rabbitMQ/logger"
)

type Config struct {
	PortEngine   string `env:"PORT_ENGINE,required"`
	APIKEY       string `env:"API_KEY,required"`
	SecretKey    string `env:"SECRET_KEY,required"`
	PublicKey    string `env:"PUBLIC_KEY,required"`
	PrivateKey   string `env:"PRIVATE_KEY,required"`
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
