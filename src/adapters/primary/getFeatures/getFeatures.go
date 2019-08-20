package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

// GetFeaturesAdapter is the interface that defines the entrypoints to this adapter
type GetFeaturesAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (Response, error)
}

type getFeaturesAdapter struct {
	service feature.FeaturePrimaryPort
}

func NewGetFeaturesAdapter(service feature.FeaturePrimaryPort) GetFeaturesAdapter {
	return &getFeaturesAdapter{
		service,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// PlaceOrder receives the request, processes it and returns a Response or an error
func (a *getFeaturesAdapter) Handle(request events.APIGatewayProxyRequest) (Response, error) {
	result := a.service.GetAll()
	fmt.Println("result:", result)

	body, err := json.Marshal(result)

	if err != nil {
		return Response{StatusCode: 400}, err
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	
	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Methods":     "GET",
			"Access-Control-Allow-Headers":     "application/json",
		},
	}

	return resp, nil
}
