package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lucasrosa/gorkin/src/adapters/secondary/database"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	repo := dynamo.NewDynamoFeatureRepository()
	service := feature.NewService(repo)
	getFeaturesAdapter := NewGetFeaturesAdapter(service)

	return getFeaturesAdapter.Handle(request)
}

func main() {
	lambda.Start(Handler)
}
