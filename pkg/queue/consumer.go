package queue

import (
	. "github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue/consumers"
)

func StartQueueConsumer() {
	Log.Debugw("Initialising RabbitMQ Main Listener ")

	conn, err := Dial()

	ch, err := conn.Channel()
	if err != nil {
		Log.Errorw("RabbitMQ Consumer Connection Error ", "error", err)
	}
	//Declare Queues
	DeclareQueues(ch)
	defer ch.Close()

	if err != nil {
		Log.Errorw("RabbitMQ Consumer Close channel Error ", "error", err)
	}

	/**
	*
	* BEGIN LISTENINNG ON QUEUES
	 */
	Log.Debugw("RabbitMQ Consumer Begin listening on queues ")
	consumers.ProcessSTKPushRequest(ch)
	consumers.ProcessSTKCallbackRequest(ch)
	/**
	*
	* BEGIN LISTENINNG ON QUEUES
	 */

	forever := make(chan bool)

	Log.Debugw("RabbitMQ Consumer Successfully Connected to our RabbitMQ Instance")
	Log.Debugw(" [*] - Waiting for messages")
	<-forever
}
