package main

import (
	"github.com/aws/aws-lambda-go/events"
	"todo-go/model"
)

// Handler POST: /todo
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	requestPath, err := parseRequestBody(request.Body)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	err = model.Put(requestBody.Id, requestBody.Title)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	return generateResponse()
}
