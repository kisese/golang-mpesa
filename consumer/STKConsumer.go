package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"gostk/config"
)

func main() {

	conn, err := amqp.Dial("amqp://" +
		"" + config.GetConfig(config.AMQP_USERNAME) + ":" +
		"" + config.GetConfig(config.AMQP_PASSWORD) + "@" +
		"" + config.GetConfig(config.AMQP_VHOST) + ":" +
		"" + config.GetConfig(config.AMQP_PORT))

	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume(
		config.GetConfig(config.STK_REQUESTS_QUEUE),
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
}
