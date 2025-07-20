package rabbitMq

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"testMod/config"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Queue      amqp.Queue
}

func Connect(config config.Config) (RabbitMQ, error) {

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config.Rabbitmq.RabbitmqUserName,
		config.Rabbitmq.RabbitmqPassword,
		config.Rabbitmq.RabbitmqHost,
		config.Rabbitmq.RabbitmqPort))
	if err != nil {
		return RabbitMQ{}, err

	}
	ch, err := conn.Channel()
	if err != nil {
		return RabbitMQ{}, err
	}

	return RabbitMQ{
		Connection: conn,
		Channel:    ch,
	}, err
}

func (r RabbitMQ) DeclareQueue(queueName string) (amqp.Queue, error) {
	queue, err := r.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return amqp.Queue{}, err
	}
	r.Queue = queue
	return queue, nil
}

func (r RabbitMQ) PublishMessage(queueName string, message string) error {
	err := r.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err

}

func (r RabbitMQ) ConsumeMessages(queueName string) (<-chan amqp.Delivery, error) {
	msgs, err := r.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
