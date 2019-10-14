package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
	"github.com/lucasrosa/gorkin/src/utils/apigateway"
)

// GetFilesAdapter is the interface that defines the entrypoints to this adapter
type GetFilesAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (apigateway.Response, error)
}

type getFilesAdapter struct {
	service feature.FilesPrimaryPort
}

// NewGetFilesAdapter returns a folder adapter with its main service (primary port)
func NewGetFilesAdapter(service feature.FilesPrimaryPort) GetFilesAdapter {
	return &getFilesAdapter{
		service,
	}
}

func (a *getFilesAdapter) Handle(request events.APIGatewayProxyRequest) (apigateway.Response, error) {
	key := request.QueryStringParameters["key"]

	result, err := a.service.Get(key)

	if err != nil {
		return apigateway.Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(map[string]interface{}{
		"url": result,
	})

	if err != nil {
		return apigateway.Response{StatusCode: 500}, err
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := apigateway.NewResponse(buf)

	return resp, nil
}
