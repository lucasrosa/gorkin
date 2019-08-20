package dynamo

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

type featureRepository struct{}

// NewDynamoFeatureRepository instantiates the repository for this adapter
func NewDynamoFeatureRepository() feature.DatabaseSecondaryPort {
	return &featureRepository{}
}

// PersistedFeature represents the model for inserting the Feature into the database
type PersistedFeature struct {
	ID        	string  	`json:"id"`
	Name 		string 		`json:"name"`
	Scenarios 	[]string 	`json:"scenarios"`
}

func (r *featureRepository) GetAll() []feature.Feature {
	fmt.Println("Database GetAll started")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}
	
	obj := []feature.Feature{}

	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("failed to make Query API call", err)
		return obj
	} 

	
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &obj)
	if err != nil {
		fmt.Println("failed to unmarshal Query result items", err)
		return obj
	}

	fmt.Println("Database GetAll ended")
	return obj
}
