package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/http/stk/forms"
	"github.com/kisese/golang_mpesa/pkg/http/stk/models"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"net/http"
	"os"
	"time"
)

type STKRequestService struct {
}

func NewStkRequestService() STKRequestService {
	return STKRequestService{}
}

func (mpesa STKRequestService) ProcessSTKPush(input forms.STKRequest) {

	timestamp := time.Now().Format("20060102150405")
	password := b64.StdEncoding.EncodeToString([]byte(
		os.Getenv("PAYBILL") +
			os.Getenv("DARAJA_PASSKEY") +
			timestamp))

	requestBody := models.STKRequestPayload{
		BusinessShortCode: os.Getenv("PAYBILL"),
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		PartyA:            input.Msisdn,
		PartyB:            os.Getenv("PAYBILL"),
		PhoneNumber:       input.Msisdn,
		AccountReference:  "DARAJA_XXX",
		TransactionDesc:   "Payment",
		Amount:            input.Amount,
		Msisdn:            input.Msisdn,
		Paybill:           os.Getenv("PAYBILL"),
		CallBackURL:       os.Getenv("STK_CALLBACK_URL"),
	}

	payload, _ := json.Marshal(requestBody)
	PostRequest(
		payload,
		getHeader(getDarajaToken()),
		os.Getenv("DARAJA_STK_URL"))
}

func getDarajaToken() interface{} {
	credential := b64.StdEncoding.EncodeToString([]byte(
		os.Getenv("DARAJA_CONSUMER_KEY") +
			":" +
			os.Getenv("DARAJA_CONSUMER_SECRET")))

	client := http.Client{}
	req, err := http.NewRequest("GET",
		os.Getenv("DARAJA_TOKEN_URL"),
		nil)

	if err != nil {
		infrastructure.Log.Errorw("Get token HTTP Error ", "error", err)
	}

	req.Header = http.Header{
		"Authorization": []string{"Basic " + credential},
	}

	resp, err := client.Do(req)
	if err != nil {
		infrastructure.Log.Fatalw("Get token HTTP Do Error ", "error", err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	infrastructure.Log.Debugw("Daraja Token Response ", "response", result, "token", result["access_token"])
	return result["access_token"]
}

func getHeader(token interface{}) map[string][]string {
	return http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + fmt.Sprintf("%v", token)},
		"cache-control": []string{"no-cache"},
	}
}
