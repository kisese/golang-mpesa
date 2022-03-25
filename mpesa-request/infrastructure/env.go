package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
)

//LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}
