package main

import (
	"encoding/json"
)

// RequestPath structure of API Request Path Parameters
type RequestPath struct {
	Id    int    `json:"id"`
}

// parseRequestBody
func parseRequestBody(body string) (RequestBody, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(body), &requestBody)

	return requestBody, err
}
