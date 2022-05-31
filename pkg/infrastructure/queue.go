package infrastructure

import (
	"github.com/streadway/amqp"
	"os"
)

func Dial() (conn *amqp.Connection, err error) {
	Log.Debugw("Initialising RabbitMQ Main Listener ")

	conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		Log.Errorw("RabbitMQ Consumer Fcailed Initializing Broker Connection ", "error", err)
		panic(err)
	}

	return
}

func DeclareQueues(ch *amqp.Channel) {
	queues := [5]string{
		os.Getenv("STK_PUSH_REQUESTS_QUEUE"),
		os.Getenv("STK_PUSH_CALLBACKS_QUEUE"),
	}

	for _, val := range queues {
		queueDeclare, err := ch.QueueDeclare(val, false, false, false, false, nil)
		Log.Debugw("RabbitMQ Publish Queue Status ", "status", queueDeclare, "queue", val)

		if err != nil {
			Log.Debugw("RabbitMQ Publish Queue Declare Error ", "err", err)
		}
	}
}
