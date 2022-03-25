package utils

import (
	b64 "encoding/base64"
	"encoding/json"
	"gostk/mpesa-request/infrastructure"
	"net/http"
	"os"
)

func GetDarajaToken() interface{} {
	darajaConsumerKey := os.Getenv("DARAJA_CONSUMER_KEY")
	darajaConsumerSecret := os.Getenv("DARAJA_CONSUMER_SECRET")
	credential := b64.StdEncoding.EncodeToString([]byte(darajaConsumerKey + ":" + darajaConsumerSecret))
	darajaTokenUrl := os.Getenv("DARAJA_TOKEN_URL")

	client := http.Client{}
	req, err := http.NewRequest("GET", darajaTokenUrl, nil)
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
