package main

import (
	"github.com/gin-gonic/gin"
	"gostk/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/stk-request", controllers.ProcessSTKPush)
	router.POST("/stk-callback", controllers.ProcessSTKCallback)

	router.Run()
}
