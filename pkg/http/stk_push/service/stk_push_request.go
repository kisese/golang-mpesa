package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/forms"
	"github.com/kisese/golang_mpesa/pkg/http/stk_push/models"
	"net/http"
	"os"
	"time"
)

type STKRequestService struct {
	input forms.STKRequest
}

func NewStkRequestService(input forms.STKRequest) STKRequestService {
	return STKRequestService{
		input: input,
	}
}

func (mpesa STKRequestService) ProcessSTKPush() {

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
		PartyA:            mpesa.input.Msisdn,
		PartyB:            os.Getenv("PAYBILL"),
		PhoneNumber:       mpesa.input.Msisdn,
		AccountReference:  "DARAJA_XXX",
		TransactionDesc:   "Payment",
		Amount:            mpesa.input.Amount,
		Msisdn:            mpesa.input.Msisdn,
		Paybill:           os.Getenv("PAYBILL"),
		CallBackURL:       os.Getenv("STK_CALLBACK_URL"),
	}

	payload, _ := json.Marshal(requestBody)
	PostRequest(
		payload,
		getHeader(getDarajaToken()),
		os.Getenv("DARAJA_STK_URL"))
}

func getHeader(token interface{}) map[string][]string {
	return http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + fmt.Sprintf("%v", token)},
		"cache-control": []string{"no-cache"},
	}
}
