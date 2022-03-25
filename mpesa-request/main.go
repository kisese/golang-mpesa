package main

import (
	"github.com/gin-gonic/gin"
	"gostk/mpesa-request/controllers"
	"gostk/mpesa-request/infrastructure"
)

func main() {
	infrastructure.LoadEnv()
	router := gin.Default()

	router.POST("/stk-request", controllers.ProcessSTKPush)

	router.Run(":8000")
}
