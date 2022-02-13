package utils

import (
	b64 "encoding/base64"
	"encoding/json"
	"gostk/logger"
	"net/http"
)

func GetDarajaToken() interface{} {
	credential := b64.StdEncoding.EncodeToString([]byte(DARAJA_CONSUMER_KEY + ":" + DARAJA_CONSUMER_SECRET))

	client := http.Client{}
	req, err := http.NewRequest("GET", DARAJA_TOKEN_URL, nil)
	if err != nil {
		logger.Log.Errorw("Get token HTTP Error ", "error", err)
	}

	req.Header = http.Header{
		"Authorization": []string{"Basic " + credential},
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.Log.Fatalw("Get token HTTP Do Error ", "error", err)
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	logger.Log.Debugw("Daraja Token Response ", "response", result, "token", result["access_token"])

	return result["access_token"]
}
