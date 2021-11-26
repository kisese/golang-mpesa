package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gostk/controllers"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/stk-request", controllers.InitiateSTKPush)

	router.Run()
}


func DotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}