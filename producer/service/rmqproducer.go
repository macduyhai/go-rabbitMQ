package service

import (
	"sync"

	"github.com/macduyhai/go-rabbitMQ/logger"
	"github.com/macduyhai/go-rabbitMQ/rabbitmq"
	"github.com/streadway/amqp"
)

var (
	one sync.Once
	rb  *RabbitMQ
)

type RMQProducerService struct {
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
	conn, errConn := rabbitmq.DialConfig(connectionString, amqp.Config{Properties: amqp.Table{"connection_name": "ProducerApp-1"}})
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

func GetRMQProducerService(rb *RabbitMQ) *RMQProducerService {
	return &RMQProducerService{
		rabbitmq: *rb,
	}
}
func (rmqservice *RMQProducerService) PublishMessage(contentType string, body []byte) error {
	// timeStart := time.Now()
	rab := rmqservice.rabbitmq
	err := rab.clientch.Publish(
		"",             // exchange
		rab.Queue.Name, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType:  contentType,
			Body:         body,
			DeliveryMode: 2,
		})
	if err != nil {
		logger.LogError("Failed to publish a message:" + err.Error())
		rab.clientCon.Close()
		return err
	}
	// logger.LogInfor(time.Since(timeStart))
	return nil
}
