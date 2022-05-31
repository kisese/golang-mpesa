package queue

import (
	"encoding/json"
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/streadway/amqp"
	"strconv"
)

func Publish(input interface{}, queue string) bool {
	//init rabbitmq
	status := true
	infrastructure.Log.Debugw("Initialising RabbitMQ Publish ", "queue", queue, "payload", input)

	conn, err := infrastructure.Dial()

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		status = false
		infrastructure.Log.Errorw("RabbitMQ Publish Connection Error ", "error", err, "queue", queue)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	queueDeclare, err := ch.QueueDeclare(
		queue,
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	infrastructure.Log.Debugw("RabbitMQ Publish Queue Status ", "status", queueDeclare, "queue", queue, "payload", input)

	// Handle any errors if we were unable to create the queue
	if err != nil {
		status = false
		fmt.Println(err)
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		status = false
		infrastructure.Log.Errorw("RabbitMQ Publish Payload Marshal Error ", "error", err, "queue", queue)
		return status
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        inputBytes,
		},
	)

	if err != nil {
		status = false
		infrastructure.Log.Errorw("RabbitMQ Publish Payload Publish Error ", "error", err, "queue", queue)
		fmt.Println(err)
	}

	infrastructure.Log.Debugw("RabbitMQ Publish Successfully Published Message to Queue ~> "+strconv.FormatBool(status), "queue", queue, "payload", input)
	return status
}
