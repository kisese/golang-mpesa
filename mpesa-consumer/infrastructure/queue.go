package infrastructure

import (
	"github.com/streadway/amqp"
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
		Log.Errorw("RabbitMQ Consumer queue consumer Error ", "error", err, "queue", queue)
	}
	return msgs
}
