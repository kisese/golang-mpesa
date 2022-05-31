package queue

import (
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue/consumers"
)

func StartQueueConsumer() {
	infrastructure.Log.Debugw("Initialising RabbitMQ Main Listener ")

	conn, err := infrastructure.Dial()

	ch, err := conn.Channel()
	if err != nil {
		infrastructure.Log.Errorw("RabbitMQ Consumer Connection Error ", "error", err)
	}
	//Declare Queues
	infrastructure.DeclareQueues(ch)
	defer ch.Close()

	if err != nil {
		infrastructure.Log.Errorw("RabbitMQ Consumer Close channel Error ", "error", err)
	}

	/**
	*
	* BEGIN LISTENINNG ON QUEUES
	 */
	infrastructure.Log.Debugw("RabbitMQ Consumer Begin listening on queues ")
	consumers.ProcessSTKPushRequest(ch)

	/**
	*
	* BEGIN LISTENINNG ON QUEUES
	 */

	forever := make(chan bool)

	infrastructure.Log.Debugw("RabbitMQ Consumer Successfully Connected to our RabbitMQ Instance")
	infrastructure.Log.Debugw(" [*] - Waiting for messages")
	<-forever
}
