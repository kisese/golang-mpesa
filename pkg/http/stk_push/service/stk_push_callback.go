package service

import (
	"encoding/json"
	. "github.com/kisese/golang_mpesa/pkg/infrastructure"
	"github.com/tidwall/gjson"
)

func ProcessSTKCallback(request map[string]interface{}) {
	bodyString, _ := json.Marshal(request)
	body := string(bodyString)

	CheckoutRequestID := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")

	Log.Debugw("Request json", "request", body, "CheckoutRequestID", CheckoutRequestID,
		"len", len(CheckoutRequestID.Str))

	if len(CheckoutRequestID.Str) > 0 {
		//Queue callback

		CheckoutRequestID := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")

		Log.Debugw("Request json", "request", body, "CheckoutRequestID", CheckoutRequestID,
			"len", len(CheckoutRequestID.Str))

		//Queue callback
		//responseCode := gjson.Get(message, "Body.stkCallback.ResultCode")
		CallbackMetadata := gjson.Get(body, "Body.stkCallback.CallbackMetadata")

		if CallbackMetadata.Exists() {
			Log.Debugw("CallbackMetadata exists")

			amount := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.0.Value")
			reference := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.1.Value")

			msisdn := gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.4.Value")
			if !msisdn.Exists() {
				msisdn = gjson.Get(body, "Body.stkCallback.CallbackMetadata.Item.3.Value")
			}

			transactionId := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")
			success := "1"
			reason := gjson.Get(body, "Body.stkCallback.ResultDesc")

			Log.Debugw("CallbackMetadata Successful Payment decoded",
				"amount", amount,
				"reference", reference,
				"msisdn", msisdn,
				"transactionId", transactionId,
				"success", success,
				"reason", reason,
			)

			//TODO Process MPESA successful STK payment
		} else {
			Log.Debugw("CallbackMetadata parsing error ", "callback", body)
			amount := ""
			reference := ""
			msisdn := ""
			transactionId := gjson.Get(body, "Body.stkCallback.CheckoutRequestID")
			success := "0"
			reason := gjson.Get(body, "Body.stkCallback.ResultDesc")

			Log.Debugw("CallbackMetadata Failed! Payment decoded",
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
		Log.Errorw("STK Callback Parse Error")
	}
}
