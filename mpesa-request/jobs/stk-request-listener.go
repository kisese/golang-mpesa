package jobs

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"gostk/mpesa-request/infrastructure"
	utils2 "gostk/mpesa-request/utils"
	"net/http"
	"os"
	"time"
)

func STKRequestListener(ch *amqp.Channel) {
	stkRequests := os.Getenv("STK_REQUESTS_QUEUE")
	infrastructure.Log.Debugw("Initialising RabbitMQ Main Listener ", "queue", stkRequests)

	queue := stkRequests
	msgs := infrastructure.Consume(ch, queue)

	go func() {
		for d := range msgs {
			var payload STKRequestPayload
			message := fmt.Sprintf("%s\n", d.Body)

			infrastructure.Log.Debugw("RabbitMQ Consumer Received Message " + message)

			//Get Payload struct from Queue
			err := json.Unmarshal([]byte(message), &payload)
			if err != nil {
				infrastructure.Log.Errorw("Payload unmarshall error ", "error", err, "queue", queue)
			}

			timestamp := time.Now().Format("20060102150405")
			callBackURL := "https://255a5142af515cb16053282b5d564739.m.pipedream.net"
			darajaPasskey := os.Getenv("DARAJA_PASSKEY")
			password := b64.StdEncoding.EncodeToString([]byte(payload.Paybill + darajaPasskey + timestamp))

			fmt.Printf("Operation: %s", payload.Msisdn)
			token := utils2.GetDarajaToken()

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

			darajaStkUrl := os.Getenv("DARAJA_STK_URL")
			utils2.PostRequest(requestBody, header, darajaStkUrl)
		}
	}()
}
