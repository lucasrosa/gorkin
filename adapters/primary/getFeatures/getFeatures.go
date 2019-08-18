package main

import (
	"github.com/aws/aws-lambda-go/events"
)

// GetFeaturesAdapter is the interface that defines the entrypoints to this adapter
type GetFeaturesAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (Response, error)
}

type getFeaturesAdapter struct {
	
}

func NewGetFeaturesAdapter() GetFeaturesAdapter {
	return &getFeaturesAdapter{}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *getFeaturesAdapter) Handle(request events.APIGatewayProxyRequest) (Response, error) {

	return successfulResponse(), nil
}

func successfulResponse() Response {
	return Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "POST",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}
}
