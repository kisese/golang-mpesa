package service

import (
	b64 "encoding/base64"
	"encoding/json"
	"github.com/kisese/golang_mpesa/pkg/infrastructure"
	"net/http"
	"os"
)

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
