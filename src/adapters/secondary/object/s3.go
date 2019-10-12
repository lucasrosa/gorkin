package s3

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

type s3i interface {
	ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
	GetObjectRequest(input *s3.GetObjectInput) (req *request.Request, output *s3.GetObjectOutput)
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

func (r *folderRepository) ListAllObjects() (feature.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(os.Getenv("BUCKET_NAME")),
		MaxKeys: aws.Int64(1000),
	}

	result, err := r.awss3.ListObjectsV2(input)
	if err != nil {
		fmt.Println("Error while trying to list objects from S3:", err.Error())
		return feature.Object{}, err
	}
	treeObject := convertToTreeObject(result.Contents)

	return treeObject, nil
}

func (r *folderRepository) GetObjectTemporaryURL(id string) (string, error) {
	req, _ := r.awss3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("BUCKET_NAME")),
		Key:    aws.String(id),
	})

	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request for object", id, "error:", err)
	}

	return urlStr, nil
}

func addChild(rootObject *feature.Object, id string, child string, grandchildren []string) {
	if rootObject.Children == nil {
		rootObject.Children = make(map[string]*feature.Object)
	}

	if len(grandchildren) == 0 {
		if child != "" { // If the string is a folder, it will end in "/", and the last item in the split will be an empty string
			if _, ok := rootObject.Children[child]; !ok { // Checks if this position in the array exists
				rootObject.Children[child] = &feature.Object{
					ID:   id,
					Name: child,
					Type: "file",
				}
			} else if rootObject.Children[child].ID == "" {
				rootObject.Children[child].ID = id
			}
		}
	} else {
		if _, ok := rootObject.Children[child]; !ok { // Checks if this position in the array exists
			rootObject.Children[child] = &feature.Object{
				ID:   id,
				Name: child,
				Type: "folder",
			}
		} else if rootObject.Children[child].ID == "" {
			rootObject.Children[child].ID = id
		}

		addChild(rootObject.Children[child], id, grandchildren[:1][0], grandchildren[1:])
	}
}

func convertToTreeObject(objects []*s3.Object) feature.Object {
	rootObject := feature.Object{ID: "root", Name: "root", Type: "folder"}
	for _, s3Object := range objects {
		etag := strings.ReplaceAll(*s3Object.ETag, "\"", "")
		if strings.Contains(*s3Object.Key, "/") {
			components := strings.Split(*s3Object.Key, "/")
			addChild(&rootObject, etag, components[0], components[1:])
		} else {
			rootObject.Children[*s3Object.Key] = &feature.Object{
				ID:   etag,
				Name: *s3Object.Key,
				Type: "file",
			}
		}
	}

	return rootObject
}
