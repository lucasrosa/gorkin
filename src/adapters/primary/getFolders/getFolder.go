package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

// GetFoldersAdapter is the interface that defines the entrypoints to this adapter
type GetFoldersAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (Response, error)
}

type getFoldersAdapter struct {
	service feature.FolderPrimaryPort
}

func NewGetFoldersAdapter(service feature.FolderPrimaryPort) GetFoldersAdapter {
	return &getFoldersAdapter{
		service,
	}
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

func (a *getFoldersAdapter) Handle(request events.APIGatewayProxyRequest) (Response, error) {
	result, err := a.service.GetAll()

	if err != nil {
		return Response{StatusCode: 500}, err
	}

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
