package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lucasrosa/gorkin/src/adapters/secondary/object"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	repo := s3.NewS3FolderRepository()
	service := feature.NewFolderService(repo)
	getFoldersAdapter := NewGetFoldersAdapter(service)

	return getFoldersAdapter.Handle(request)
}

func main() {
	lambda.Start(Handler)
}
