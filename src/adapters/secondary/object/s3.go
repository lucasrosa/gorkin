package s3

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

type s3i interface {
	ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
}

type folderRepository struct {
	awss3 s3i
}

// NewS3FolderRepository instantiates the repository for this adapter
func NewS3FolderRepository() (feature.ObjectSecondaryPort, error) {
	svc, err := newS3Session()

	if err != nil {
		return &folderRepository{}, err
	}

	return &folderRepository{
		svc,
	}, nil
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

// listChildren just gets the CommonPrefixes (folders) and Contents (files)
// and returns them as a list of strings
func listChildren(objects *s3.ListObjectsV2Output, folderName string) []string {
	var children []string

	// Get all folders
	for _, doc := range objects.CommonPrefixes {
		children = append(children, *doc.Prefix)
	}

	// Get all files
	for _, doc := range objects.Contents {
		// S3 lists the folder itself, so we have to remove it from the list
		if *doc.Key != folderName {
			children = append(children, *doc.Key)
		}
	}

	return children
}

// ListObjects lists all items inside a given folder in AWS S3
func (r *folderRepository) ListObjects(folder string) (feature.Folder, error) {

	input := &s3.ListObjectsV2Input{
		Bucket:    aws.String(os.Getenv("BUCKET_NAME")),
		MaxKeys:   aws.Int64(128),
		Delimiter: aws.String("/"),
	}

	// If there is no folder, leave the Prefix empty so S3 lists the root folders
	if folder != "" {
		input.SetPrefix(folder)
	}

	result, err := r.awss3.ListObjectsV2(input)
	if err != nil {
		fmt.Println("Error while trying to list objects from S3:", err.Error())
		return feature.Folder{}, err
	}

	children := listChildren(result, folder)

	myFolder := feature.Folder{
		Children: children,
	}

	return myFolder, nil
}
