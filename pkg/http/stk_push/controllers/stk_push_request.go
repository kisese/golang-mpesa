package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/forms"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/kisese/golang_mpesa/pkg/queue"
	"github.com/kisese/golang_mpesa/pkg/utils"
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

	queue.Publish(request, os.Getenv("STK_PUSH_CALLBACKS_QUEUE"))

	context.JSON(http.StatusOK, gin.H{"ok": "ok"})
}
