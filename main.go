package main

import (
	"github.com/gin-gonic/gin"
	"gostk/controllers"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	router.POST("/stk-request", controllers.InitiateSTKPush)

	router.Run()
}
