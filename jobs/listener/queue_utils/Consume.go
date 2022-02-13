package queue_utils

import (
	"github.com/streadway/amqp"
	"gostk/logger"
)

func Consume(ch *amqp.Channel, queue string) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		logger.Log.Errorw("RabbitMQ Consumer queue consumer Error ", "error", err, "queue", queue)
	}
	return msgs
}
