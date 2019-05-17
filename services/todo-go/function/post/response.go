package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response
type Response events.APIGatewayProxyResponse

// ResponseBody
type ResponseBody struct {
	Message string `json:"message"`
}

// generateResponse
func generateResponse() (Response, error) {
	responseBody := ResponseBody{
		Message: "PutItem Succeeded.",
	}

	body, err := json.Marshal(responseBody)

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	headers := map[string]string{
		"Content-Type":           "application/json",
		"X-MyCompany-Func-Reply": "hello-handler",
	}

	response := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(body),
		Headers:         headers,
	}

	return response, nil
}
