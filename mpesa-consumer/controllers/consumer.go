package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gostk/mpesa-consumer/infrastructure"
	"gostk/mpesa-consumer/jobs"
	"io/ioutil"
	"net/http"
	"os"
)

func ProcessSTKCallback(context *gin.Context) {
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
		stkCallbacks := os.Getenv("STK_CALLBACKS")
		jobs.Publish(request, stkCallbacks)
	} else {
		infrastructure.Log.Errorw("STK Callback Parse Error")
		context.JSON(http.StatusBadRequest, gin.H{"error": "STK Callback Parse Error"})
	}

	context.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
