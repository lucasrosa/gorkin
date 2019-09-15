package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

type folderRepository struct{}

// NewS3FolderRepository instantiates the repository for this adapter
func NewS3FolderRepository() feature.ObjectSecondaryPort {
	return &folderRepository{}
}

func (r *folderRepository) GetAll(folder string) []feature.Folder {
	fmt.Println("Folder GetAll started, looking into bucket:", os.Getenv("BUCKET_NAME"))

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error while creating session")
	}

	svc := s3.New(sess)

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("gorkin-features-dev"),
		//Bucket:  aws.String(os.Getenv("BUCKET_NAME")),
		MaxKeys:   aws.Int64(128),
		Delimiter: aws.String("/"),
	}

	if folder != "" {
		input.SetPrefix(folder)
		fmt.Println("querying folder", folder)
	}

	var children []string

	result, err := svc.ListObjectsV2(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fmt.Println("1:", s3.ErrCodeNoSuchBucket, aerr.Error())
			default:
				fmt.Println("2:", aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println("3:", err.Error())
		}
		fmt.Println("NOK")
	} else {
		fmt.Println("OK")
		for index, doc := range result.CommonPrefixes {
			fmt.Println("Folder", index)
			children = append(children, *doc.Prefix)
			fmt.Println(doc)
		}

		for index, doc := range result.Contents {
			fmt.Println("Content", index)
			fmt.Println(doc)
			if *doc.Key != folder {
				children = append(children, *doc.Key)
			}
		}
	}

	myFolder := feature.Folder{
		Children: children,
	}
	obj := []feature.Folder{}
	obj = append(obj, myFolder)

	return obj
}
