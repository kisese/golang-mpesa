package jobs

import (
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tidwall/gjson"
	"gostk/mpesa-consumer/infrastructure"
	"os"
)

func STKCallbackListener(ch *amqp.Channel) {
	stkCallbacks := os.Getenv("STK_CALLBACKS_QUEUE")

	infrastructure.Log.Debugw("Initialising RabbitMQ Main Listener ", "queue", stkCallbacks)

	queue := stkCallbacks
	msgs := infrastructure.Consume(ch, queue)

	go func() {
		for d := range msgs {
			message := fmt.Sprintf("%s\n", d.Body)

			infrastructure.Log.Debugw("RabbitMQ Consumer Received Message " + message)

			CheckoutRequestID := gjson.Get(message, "Body.stkCallback.CheckoutRequestID")
			infrastructure.Log.Debugw("Request json", "request", message, "CheckoutRequestID", CheckoutRequestID,
				"len", len(CheckoutRequestID.Str))

			if len(message) > 0 {
				//Queue callback
				//responseCode := gjson.Get(message, "Body.stkCallback.ResultCode")
				CallbackMetadata := gjson.Get(message, "Body.stkCallback.CallbackMetadata")

				if CallbackMetadata.Exists() {
					infrastructure.Log.Debugw("CallbackMetadata exists")

					amount := gjson.Get(message, "Body.stkCallback.CallbackMetadata.Item.0.Value")
					reference := gjson.Get(message, "Body.stkCallback.CallbackMetadata.Item.1.Value")

					msisdn := gjson.Get(message, "Body.stkCallback.CallbackMetadata.Item.4.Value")
					if !msisdn.Exists() {
						msisdn = gjson.Get(message, "Body.stkCallback.CallbackMetadata.Item.3.Value")
					}

					transactionId := gjson.Get(message, "Body.stkCallback.CheckoutRequestID")
					success := "1"
					reason := gjson.Get(message, "Body.stkCallback.ResultDesc")

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
					infrastructure.Log.Debugw("CallbackMetadata parsing error ", "callback", message)
					amount := ""
					reference := ""
					msisdn := ""
					transactionId := gjson.Get(message, "Body.stkCallback.CheckoutRequestID")
					success := "0"
					reason := gjson.Get(message, "Body.stkCallback.ResultDesc")

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
			}

		}
	}()
}
