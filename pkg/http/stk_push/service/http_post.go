package service

import (
	"bytes"
	"encoding/json"
	. "github.com/kisese/golang_mpesa/pkg/infrastructure"
	"net/http"
)

func PostRequest(requestBody []byte, header http.Header, url string) interface{} {

	Log.Debugw("PostRequest() ", "header", header)

	client := http.Client{}

	responseBody := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", url, responseBody)
	if err != nil {
		Log.Errorw("POST PostRequest HTTP Error ", "error", err, "url", url)
	}

	req.Header = header
	resp, err := client.Do(req)
	if err != nil {
		Log.Fatalw("POST token HTTP Do Error ", "error", err, "url", url)
	}

	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	Log.Debugw("POST Response ", "response", result, "url", url)

	return result
}
