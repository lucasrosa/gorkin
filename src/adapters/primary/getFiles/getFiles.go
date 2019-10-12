package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

// GetFilesAdapter is the interface that defines the entrypoints to this adapter
type GetFilesAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (Response, error)
}

type getFilesAdapter struct {
	service feature.FilesPrimaryPort
}

func NewGetFilesAdapter(service feature.FilesPrimaryPort) GetFilesAdapter {
	return &getFilesAdapter{
		service,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func (a *getFilesAdapter) Handle(request events.APIGatewayProxyRequest) (Response, error) {
	key := request.QueryStringParameters["key"]

	result, err := a.service.Get(key)

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"url": result,
	})

	if err != nil {
		return Response{StatusCode: 500}, err
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
