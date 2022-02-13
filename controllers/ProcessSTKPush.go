package controllers

import (
	"github.com/gin-gonic/gin"
	"gostk/jobs/publisher"
	"gostk/jobs/requests"
	"gostk/utils"
	"gostk/validation"
	"net/http"
	"os"
)

func InitiateSTKPush(context *gin.Context) {

	var input validation.STKRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stkPayload := requests.STKRequestPayload{
		Amount:          "10",
		Msisdn:          "254720000000",
		Paybill:         os.Getenv("PAYBILL"),
		TrxId:           "12345678",
		ReferenceNumber: "12345678",
		CallbackUrl:     "https://255a5142af515cb16053282b5d564739.m.pipedream.net",
	}

	publisher.Publish(stkPayload, utils.STK_REQUESTS)

	context.JSON(http.StatusOK, gin.H{"data": "Please wait to enter pin"})
}
