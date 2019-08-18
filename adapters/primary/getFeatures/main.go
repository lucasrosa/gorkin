package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	// checkoutRepo := queueadaptersqs.NewSQSCheckoutRepository()
	// checkoutService := cart.NewCheckoutService(checkoutRepo)
	getFeaturesAdapter := NewGetFeaturesAdapter()

	return getFeaturesAdapter.Handle(request)
}

func main() {
	lambda.Start(Handler)
}
