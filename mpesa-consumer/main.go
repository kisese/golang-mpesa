package main

import (
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"gostk/mpesa-consumer/controllers"
	"gostk/mpesa-consumer/infrastructure"
	"gostk/mpesa-consumer/jobs"
)

func main() {
	infrastructure.LoadEnv()
	router := gin.Default()

	router.POST("/stk-callback", controllers.ProcessSTKCallback)

	router.Run(":8001")
	InitiateQueue()
}

func InitiateQueue() {
	infrastructure.Log.Debugw("Initialising RabbitMQ Main Listener ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		infrastructure.Log.Errorw("RabbitMQ Consumer Failed Initializing Broker Connection ", "error", err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		infrastructure.Log.Errorw("RabbitMQ Consumer Connection Error ", "error", err)
	}
	defer ch.Close()

	if err != nil {
		infrastructure.Log.Errorw("RabbitMQ Consumer Close channel Error ", "error", err)
	}

	//BEGIN LISTENINNG ON QUEUES
	infrastructure.Log.Debugw("RabbitMQ Consumer Begin listening on queues ")
	jobs.STKCallbackListener(ch)

	forever := make(chan bool)

	infrastructure.Log.Debugw("RabbitMQ Consumer Successfully Connected to our RabbitMQ Instance")
	infrastructure.Log.Debugw(" [*] - Waiting for messages")
	<-forever
}
