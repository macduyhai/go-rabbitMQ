package service

import (
	"fmt"
	"sync"

	"github.com/macduyhai/go-rabbitMQ/logger"
	"github.com/macduyhai/go-rabbitMQ/rabbitmq"
	"github.com/streadway/amqp"
)

var (
	one sync.Once
	rb  *RabbitMQ
)

type RMQConsumerService struct {
	rabbitmq RabbitMQ
}
type RabbitMQ struct {
	clientCon *rabbitmq.Connection
	clientch  *rabbitmq.Channel
	Queue     *amqp.Queue
	// 	rbmqErr   error
}

func InitConnectionRB(connectionString string, queueString string) *RabbitMQ {
	logger.LogInfor("Connecting to RB ...")
	rb = &RabbitMQ{}
	conn, errConn := rabbitmq.DialConfig(connectionString, amqp.Config{Properties: amqp.Table{"connection_name": "ConsumerApp-1"}})
	if errConn != nil {
		logger.LogError("Failed to connect to RabbitMQ:" + errConn.Error())
	}
	ch, err := conn.Channel()
	if err != nil {
		logger.LogError("Failed to open a channel:" + err.Error())
	}
	q, err := ch.QueueDeclare(
		queueString, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		logger.LogError("Failed to declare a queue:" + err.Error())
	}
	rb.clientCon = conn
	rb.clientch = ch
	rb.Queue = &q

	return rb
}

func GetRMQConsumerService(rb *RabbitMQ) *RMQConsumerService {
	return &RMQConsumerService{
		rabbitmq: *rb,
	}
}

func (rmqservice *RMQConsumerService) HanderConsumer() error {
	msgs, err := rmqservice.rabbitmq.clientch.Consume(
		rmqservice.rabbitmq.Queue.Name, // queue
		"",                             // consumer
		true,                           // auto-acks
		false,                          // exclusive
		false,                          // no-local
		false,                          // no-wait
		nil,                            // args
	)
	if err != nil {
		logger.LogError("Client Consumer error")
		return err
	}

	forever := make(chan bool)
	go func() {
		index := 0
		for d := range msgs {
			index++
			logger.LogInfor(string(d.Body))
			logger.LogInfor(index)
		}

	}()
	fmt.Printf("Started listening for messages on '%s' queue\n", rmqservice.rabbitmq.Queue.Name)
	<-forever
	return nil

}
