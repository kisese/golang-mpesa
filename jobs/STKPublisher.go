package jobs

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"gostk/config"
	"gostk/requests"
)

func PublishMessage(input requests.STKRequest, queue_name string) {

	conn, err := amqp.Dial("amqp://" +
		"" + config.GetConfig(config.AMQP_USERNAME) + ":" +
		"" + config.GetConfig(config.AMQP_PASSWORD) + "@" +
		"" + config.GetConfig(config.AMQP_VHOST) + ":" +
		"" + config.GetConfig(config.AMQP_PORT))

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	// with this channel open, we can then start to interact
	// with the instance and declare Queues that we can publish and
	// subscribe to
	q, err := ch.QueueDeclare(
		queue_name,
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println(err)
	}

	inputBytes, err := json.Marshal(input)
	if err != nil {
		print(err)
		return
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		"",
		queue_name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        inputBytes,
		},
	)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Published Message to Queue")
}
