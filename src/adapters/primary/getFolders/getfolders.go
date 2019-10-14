package main

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
	"github.com/lucasrosa/gorkin/src/utils/apigateway"
)

// GetFoldersAdapter is the interface that defines the entrypoints to this adapter
type GetFoldersAdapter interface {
	Handle(request events.APIGatewayProxyRequest) (apigateway.Response, error)
}

type getFoldersAdapter struct {
	service feature.FolderPrimaryPort
}

// NewGetFoldersAdapter returns a folder adapter with its main service (primary port)
func NewGetFoldersAdapter(service feature.FolderPrimaryPort) GetFoldersAdapter {
	return &getFoldersAdapter{
		service,
	}
}

func (a *getFoldersAdapter) Handle(request events.APIGatewayProxyRequest) (apigateway.Response, error) {
	result, err := a.service.GetAll()

	if err != nil {
		return apigateway.Response{StatusCode: 500}, err
	}

	body, err := json.Marshal(result)

	if err != nil {
		return apigateway.Response{StatusCode: 400}, err
	}

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	resp := apigateway.NewResponse(buf)

	return resp, nil
}
