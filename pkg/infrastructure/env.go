package infrastructure

import (
	"github.com/joho/godotenv"
	"github.com/kisese/golang_mpesa/pkg/utils"
	"log"
)

//LoadEnv loads environment variables from .env file
func LoadEnv() {

	err := godotenv.Load(utils.GetPath() + `../pkg/.env`)

	//err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file " + err.Error())
	}
}
