package listeners

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"gostk/jobs/listener/queue_utils"
	"gostk/jobs/requests"
	"gostk/logger"
	"gostk/utils"
	"net/http"
	"time"
)

func STKRequestListener(ch *amqp.Channel) {
	logger.Log.Debugw("Initialising RabbitMQ Main Listener ", "queue", utils.STK_REQUESTS)

	queue := utils.STK_REQUESTS
	msgs := queue_utils.Consume(ch, queue)

	go func() {
		for d := range msgs {
			var payload requests.STKRequestPayload
			message := fmt.Sprintf("%s\n", d.Body)

			logger.Log.Debugw("RabbitMQ Consumer Received Message " + message)

			//Get Payload struct from Queue
			err := json.Unmarshal([]byte(message), &payload)
			if err != nil {
				logger.Log.Errorw("Payload unmarshall error ", "error", err, "queue", queue)
			}

			timestamp := time.Now().Format("20060102150405")
			callBackURL := "https://255a5142af515cb16053282b5d564739.m.pipedream.net"
			password := b64.StdEncoding.EncodeToString([]byte(payload.Paybill + utils.DARAJA_PASSKEY + timestamp))

			fmt.Printf("Operation: %s", payload.Msisdn)
			token := utils.GetDarajaToken()

			requestBody, _ := json.Marshal(map[string]string{
				"BusinessShortCode": payload.Paybill,
				"Password":          password,
				"Timestamp":         timestamp,
				"TransactionType":   "CustomerPayBillOnline",
				"Amount":            payload.Amount,
				"PartyA":            payload.Msisdn,
				"PartyB":            payload.Paybill,
				"PhoneNumber":       payload.Msisdn,
				"CallBackURL":       callBackURL,
				"AccountReference":  "DARAJA_XXX",
				"TransactionDesc":   "Payment",
			})

			header := http.Header{
				"Content-Type":  []string{"application/json"},
				"Authorization": []string{"Bearer " + fmt.Sprintf("%v", token)},
				"cache-control": []string{"no-cache"},
			}

			utils.PostRequest(requestBody, header, utils.DARAJA_STK_URL)
		}
	}()
}
