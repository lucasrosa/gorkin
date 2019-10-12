package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	s3 "github.com/lucasrosa/gorkin/src/adapters/secondary/object"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
	verifier "github.com/okta/okta-jwt-verifier-golang"
)

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (Response, error) {
	fmt.Println("request Authorization:")
	fmt.Println(request.Headers["Authorization"])

	// if !isAuthenticated(request.Headers["Authorization"]) {
	// 	return Response{StatusCode: 401}, nil
	// }
	repo, err := s3.NewS3FolderRepository()

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	service := feature.NewFolderService(repo)
	getFoldersAdapter := NewGetFoldersAdapter(service)

	return getFoldersAdapter.Handle(request)
}

func main() {
	lambda.Start(Handler)
}

func isAuthenticated(authHeader string) bool {
	//authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		return false
	}
	tokenParts := strings.Split(authHeader, "Bearer ")
	bearerToken := tokenParts[1]

	tv := map[string]string{}
	tv["aud"] = "api://default"
	tv["cid"] = "0oa1iwp6inUyX1mys357" //os.Getenv("SPA_CLIENT_ID")
	jv := verifier.JwtVerifier{
		Issuer:           "https://dev-774764.okta.com/oauth2/default", //os.Getenv("ISSUER"),
		ClaimsToValidate: tv,
	}

	_, err := jv.New().VerifyAccessToken(bearerToken)

	if err != nil {
		fmt.Println("Error while authenticating: ", err)
		return false
	}

	return true
}
