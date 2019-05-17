package main

import (
	"encoding/json"
)

// RequestBody structure of API Request Parameters
type RequestBody struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// parseRequestBody
func parseRequestBody(body string) (RequestBody, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(body), &requestBody)

	return requestBody, err
}
