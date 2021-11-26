package controllers

import (
	"github.com/gin-gonic/gin"
	"gostk/config"
	"gostk/jobs"
	"gostk/requests"
	"net/http"
)

func InitiateSTKPush(context *gin.Context) {

	var input requests.STKRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jobs.PublishMessage(input, config.GetConfig(config.STK_REQUESTS_QUEUE))

	context.JSON(http.StatusOK, gin.H{"data": "Please wait to enter pin"})
}
