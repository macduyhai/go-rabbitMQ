package service

import (
	"github.com/macduyhai/go-rabbitMQ/logger"
	"github.com/streadway/amqp"
)

type RMQProducerService struct {
	Queue     *amqp.Queue
	clientCon *amqp.Connection
	clientch  *amqp.Channel
	rbmqErr   error
}

func GetRMQProducerService(connectionString string, queueString string) *RMQProducerService {
	var rmqservice = &RMQProducerService{}
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		logger.LogError("Failed to connect to RabbitMQ:" + err.Error())
		rmqservice.rbmqErr = err
		return rmqservice
	}
	ch, err := conn.Channel()
	if err != nil {
		logger.LogError("Failed to open a channel:" + err.Error())
		rmqservice.rbmqErr = err
		return rmqservice
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
		rmqservice.rbmqErr = err
		return rmqservice
	}

	rmqservice.clientCon = conn
	rmqservice.clientch = ch
	rmqservice.Queue = &q
	rmqservice.rbmqErr = nil

	// defer conn.Close()
	return rmqservice
}
func (rmqservice *RMQProducerService) PublishMessage(contentType string, body []byte) error {
	// timeStart := time.Now()
	err := rmqservice.clientch.Publish(
		"",                    // exchange
		rmqservice.Queue.Name, // routing key
		false,                 // mandatory
		false,                 // immediate
		amqp.Publishing{
			ContentType:  contentType,
			Body:         body,
			DeliveryMode: 2,
		})
	if err != nil {
		logger.LogError("Failed to publish a message:" + err.Error())

		return err
	}
	// logger.LogInfor(time.Since(timeStart))
	return nil
}
