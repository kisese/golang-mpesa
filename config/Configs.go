package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const STK_REQUESTS_QUEUE = "STK_REQUESTS_QUEUE"
const AMQP_USERNAME = "AMQP_USERNAME"
const AMQP_PASSWORD = "AMQP_PASSWORD"
const AMQP_VHOST = "AMQP_VHOST"
const AMQP_PORT = "AMQP_PORT"

func GetConfig(key string) string {

	// load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
