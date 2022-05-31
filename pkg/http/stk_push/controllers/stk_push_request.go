package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/forms"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue"
	"github.com/kisese/golang_mpesa/pkg/utils"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
)

type STKController struct {
}

func NewSTRRequestController() STKController {
	return STKController{}
}

func (mpesa *STKController) ProcessSTKPushRequest(context *gin.Context) {

	var input forms.STKRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		infrastructure.Log.Errorw("edit profile validation error", "error", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	queue.Publish(input, os.Getenv("STK_PUSH_REQUESTS_QUEUE"))

	utils.SuccessJSON(context, http.StatusOK, "STK Push Request Received")
}

func (mpesa *STKController) ProcessSTKCallback(context *gin.Context) {
	infrastructure.Log.Debugw("Init STK Callback")

	ByteBody, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))

	var request map[string]interface{}
	json.NewDecoder(context.Request.Body).Decode(&request)
	bodyString, _ := json.Marshal(request)
	body := string(bodyString)

	CheckoutRequestID := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")

	infrastructure.Log.Debugw("Request json", "request", body, "CheckoutRequestID", CheckoutRequestID,
		"len", len(CheckoutRequestID.Str))

	if len(CheckoutRequestID.Str) > 0 {
		//Queue callback

		CheckoutRequestID := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")

		infrastructure.Log.Debugw("Request json", "request", body, "CheckoutRequestID", CheckoutRequestID,
			"len", len(CheckoutRequestID.Str))

		//Queue callback
		//responseCode := gjson.Get(message, "Body.stkCallback.ResultCode")
		CallbackMetadata := gjson.Get(body, "Body.stkCallback.CallbackMetadata")

		if CallbackMetadata.Exists() {
			infrastructure.Log.Debugw("CallbackMetadata exists")

			amount := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.0.Value")
			reference := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.1.Value")

			msisdn := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.4.Value")
			if !msisdn.Exists() {
				msisdn = gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.3.Value")
			}

			transactionId := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")
			success := "1"
			reason := gjson.Get(body, "Body.stkCallback.ResultDesc")

			infrastructure.Log.Debugw("CallbackMetadata Successful Payment decoded",
				"amount", amount,
				"reference", reference,
				"msisdn", msisdn,
				"transactionId", transactionId,
				"success", success,
				"reason", reason,
			)

			//TODO Process MPESA successful STK payment
		} else {
			infrastructure.Log.Debugw("CallbackMetadata parsing error ", "callback", body)
			amount := ""
			reference := ""
			msisdn := ""
			transactionId := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")
			success := "0"
			reason := gjson.Get(body, "Body.stkCallback.ResultDesc")

			infrastructure.Log.Debugw("CallbackMetadata Failed! Payment decoded",
				"amount", amount,
				"reference", reference,
				"msisdn", msisdn,
				"transactionId", transactionId,
				"success", success,
				"reason", reason,
			)

			//TODO Process MPESA Failed! STK payment
		}

	} else {
		infrastructure.Log.Errorw("STK Callback Parse Error")
		context.JSON(http.StatusBadRequest, gin.H{"error": "STK Callback Parse Error"})
	}

	context.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
