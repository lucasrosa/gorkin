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

func newS3Session() (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		fmt.Println("Error while creating session with AWS")
		return nil, err
	}

	return s3.New(sess), nil
}

// ListObjects lists all items inside a given folder in AWS S3
func (r *folderRepository) ListObjects(folder string) (feature.Folder, error) {
	fmt.Println("Folder ListObjects started, looking into bucket:", os.Getenv("BUCKET_NAME"))

	svc, err := newS3Session()
	if err != nil {
		return feature.Folder{}, err
	}

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(os.Getenv("BUCKET_NAME")),
		MaxKeys:   aws.Int64(128),
		Delimiter: aws.String("/"),
	}

	// If there is no folder, leave the Prefix empty so S3 lists the root folders
	if folder != "" {
		input.SetPrefix(folder)
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
			// Cast err to awserr.Error to get the Code and Message from an error.
			fmt.Println("3:", err.Error())
		}

		return feature.Folder{}, err
	}

	// Get all folders
	for _, doc := range result.CommonPrefixes {
		children = append(children, *doc.Prefix)
	}

	// Get all files
	for _, doc := range result.Contents {
		if *doc.Key != folder {
			children = append(children, *doc.Key)
		}
	}

	myFolder := feature.Folder{
		Children: children,
	}

	return myFolder, nil
}
