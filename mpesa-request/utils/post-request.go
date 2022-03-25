package utils

import (
	"bytes"
	"encoding/json"
	"gostk/mpesa-request/infrastructure"
	"net/http"
)

func PostRequest(requestBody []byte, header http.Header, url string) interface{} {

	infrastructure.Log.Debugw("PostRequest() ", "payload", requestBody, "headers", header)

	client := http.Client{}

	responseBody := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		infrastructure.Log.Errorw("POST PostRequest HTTP Error ", "error", err, "url", url)
	}

	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		infrastructure.Log.Fatalw("POST token HTTP Do Error ", "error", err, "url", url)
	}

	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	infrastructure.Log.Debugw("POST Response ", "response", result, "url", url)

	return result
}
