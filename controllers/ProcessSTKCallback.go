package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"gostk/jobs/publisher"
	"gostk/logger"
	"gostk/utils"
	"io/ioutil"
	"net/http"
)

func ProcessSTKCallback(context *gin.Context) {
	logger.Log.Debugw("Init STK Callback")

	ByteBody, _ := ioutil.ReadAll(context.Request.Body)
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))

	var request map[string]interface{}
	json.NewDecoder(context.Request.Body).Decode(&request)
	bodyString, _ := json.Marshal(request)
	body := string(bodyString)

	CheckoutRequestID := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")

	logger.Log.Debugw("Request json", "request", body, "CheckoutRequestID", CheckoutRequestID,
		"len", len(CheckoutRequestID.Str))

	if len(CheckoutRequestID.Str) > 0 {
		//Queue callback
		publisher.Publish(request, utils.STK_CALLBACKS)
	} else {
		logger.Log.Errorw("STK Callback Parse Error")
		context.JSON(http.StatusBadRequest, gin.H{"error": "STK Callback Parse Error"})
	}

	context.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
