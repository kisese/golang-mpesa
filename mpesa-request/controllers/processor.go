package controllers

import (
	"github.com/gin-gonic/gin"
	"gostk/mpesa-request/infrastructure"
	"gostk/mpesa-request/jobs"
	"gostk/mpesa-request/validation"
	"net/http"
	"os"
)

func ProcessSTKPush(context *gin.Context) {

	var input validation.STKRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		infrastructure.Log.Errorw("edit profile validation error", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stkPayload := jobs.STKRequestPayload{
		Amount:      input.Amount,
		Msisdn:      input.Msisdn,
		Paybill:     os.Getenv("PAYBILL"),
		CallbackUrl: "https://255a5142af515cb16053282b5d564739.m.pipedream.net",
	}

	stkRequests := os.Getenv("STK_REQUESTS_QUEUE")
	jobs.Publish(stkPayload, stkRequests)

	context.JSON(http.StatusOK, gin.H{"data": "Please wait to enter pin"})
}
