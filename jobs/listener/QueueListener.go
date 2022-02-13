package main

import (
	"github.com/streadway/amqp"
	"gostk/jobs/listener/listeners"
	"gostk/logger"
)

func main() {
	logger.Log.Debugw("Initialising RabbitMQ Main Listener ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.Log.Errorw("RabbitMQ Consumer Failed Initializing Broker Connection ", "error", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		logger.Log.Errorw("RabbitMQ Consumer Connection Error ", "error", err)
	}
	defer ch.Close()

	if err != nil {
		logger.Log.Errorw("RabbitMQ Consumer Close channel Error ", "error", err)
	}

	//BEGIN LISTENINNG ON QUEUES
	logger.Log.Debugw("RabbitMQ Consumer Begin listening on queues ")
	listeners.STKRequestListener(ch)
	listeners.STKCallbackListener(ch)

	forever := make(chan bool)

	logger.Log.Debugw("RabbitMQ Consumer Successfully Connected to our RabbitMQ Instance")
	logger.Log.Debugw(" [*] - Waiting for messages")
	<-forever
}
