package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	s3 "github.com/lucasrosa/gorkin/src/adapters/secondary/object"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
	"github.com/lucasrosa/gorkin/src/utils/apigateway"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (apigateway.Response, error) {

	// if !isAuthenticated(request.Headers["Authorization"]) {
	// 	return Response{StatusCode: 401}, nil
	// }
	repo, err := s3.NewS3FolderRepository()

	if err != nil {
		return apigateway.Response{StatusCode: 500}, err
	}

	service := feature.NewFilesService(repo)
	getFilesAdapter := NewGetFilesAdapter(service)

	return getFilesAdapter.Handle(request)
}

func main() {
	lambda.Start(Handler)
}
