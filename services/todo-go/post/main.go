package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// RequestBody structure of API Request Parameters
type RequestBody struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// ResponseBody structure of API Response Parameters
type ResponseBody struct {
	Message string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	requestBody, err := parseRequestBody(request.Body)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	err = putItem(requestBody.Id, requestBody.Title)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	return generateResponse()
}

// parseRequestBody
func parseRequestBody(body string) (RequestBody, error) {
	var requestBody RequestBody
	err := json.Unmarshal([]byte(body), &requestBody)

	return requestBody, err
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

func main() {
	lambda.Start(Handler)
}

// Todo structure of DynamoDB
// TDOD: Model層に寄せる
type Todo struct {
	Id        int       `dynamo:"id"`
	Title     string    `dynamo:"title"`
	CreatedAt time.Time `dynamo:"created_at"`
}

// putItem
// TODO: Model層に寄せる
func putItem(id int, title string) error {
	item := Todo{
		Id:        id,
		Title:     title,
		CreatedAt: time.Now().UTC(),
	}

	db := dynamo.New(session.New())
	table := db.Table("Todo")

	return table.Put(item).Run()
}
