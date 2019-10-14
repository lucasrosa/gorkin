package s3

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/lucasrosa/gorkin/src/corelogic/feature"
)

func TestListChildrenWithNoItems(t *testing.T) {

	outputObject := s3.ListObjectsV2Output{
		CommonPrefixes: []*s3.CommonPrefix{},
		Contents:       []*s3.Object{},
	}

	var expected []string

	got := listChildren(&outputObject, "folder0")

	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected", expected, "got", got)
	}

}

type s3Mock struct{}

func (s *s3Mock) GetObjectRequest(input *s3.GetObjectInput) (req *request.Request, output *s3.GetObjectOutput) {
	return &request.Request{}, &s3.GetObjectOutput{}
}

// ListObjectsV2 returns an object with 2 folders and 2 files
func (s *s3Mock) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	time0 := time.Now()
	var size0 int64 = 0
	storage_class0 := "STANDARD"
	// Folders 1, 2 and 3 have the same ETag
	etag1 := "\"d41d8cd98f00b204e9800998ecf8427e\""
	key1 := "folder1/"
	key2 := "folder1/folder1_1/"
	key3 := "folder1/folder1_1/folder1_1_1/"

	etag4 := "\"73a82eb7b09a2c5499cc76c2aefe52e2\""
	key4 := "folder1/folder1_1/folder1_1_1/12362714-dzone-refcard215-microservices.pdf"
	var size4 int64 = 408400

	etag5 := "\"9b1529ddfd06b2046b2615f58ad2829f\""
	key5 := "folder1/folder1_1/folder1_1_1/halo.txt"
	var size5 int64 = 6

	// Folders 6 and 8 have the same ETag
	etag6 := "\"d41d8cd98f00b204e9800998ecf8427e\""
	key6 := "folder2/"

	etag7 := "\"5053efd24093110335a9a3c3c6dd17f8\""
	key7 := "folder2/BP-Diet-Roadmap-2019.pdf"
	var size7 int64 = 1487993

	key8 := "folder2/folder2_1/"

	etag9 := "\"1ca5cf539336c45b48ae2369a56b40bf\""
	key9 := "foto.png"
	var size9 int64 = 219555

	outputObject := s3.ListObjectsV2Output{
		Contents: []*s3.Object{
			&s3.Object{
				ETag:         &etag1,
				Key:          &key1,
				LastModified: &time0,
				Size:         &size0,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag1,
				Key:          &key2,
				LastModified: &time0,
				Size:         &size0,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag1,
				Key:          &key3,
				LastModified: &time0,
				Size:         &size0,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag4,
				Key:          &key4,
				LastModified: &time0,
				Size:         &size4,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag5,
				Key:          &key5,
				LastModified: &time0,
				Size:         &size5,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag6,
				Key:          &key6,
				LastModified: &time0,
				Size:         &size0,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag7,
				Key:          &key7,
				LastModified: &time0,
				Size:         &size7,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag6,
				Key:          &key8,
				LastModified: &time0,
				Size:         &size0,
				StorageClass: &storage_class0,
			},
			&s3.Object{
				ETag:         &etag9,
				Key:          &key9,
				LastModified: &time0,
				Size:         &size9,
				StorageClass: &storage_class0,
			},
		},
	}

	return &outputObject, nil
}

func TestListAllObjects(t *testing.T) {
	expected := feature.Object{
		Key:  "root",
		Name: "root",
		Type: "folder",
		Children: map[string]*feature.Object{
			"folder1": &feature.Object{
				Key:  "folder1/",
				Name: "folder1",
				Type: "folder",
				Children: map[string]*feature.Object{
					"folder1_1": &feature.Object{
						Key:  "folder1/folder1_1/",
						Name: "folder1_1",
						Type: "folder",
						Children: map[string]*feature.Object{
							"folder1_1_1": &feature.Object{
								Key:  "folder1/folder1_1/folder1_1_1/",
								Name: "folder1_1_1",
								Type: "folder",
								Children: map[string]*feature.Object{
									"12362714-dzone-refcard215-microservices.pdf": &feature.Object{
										Key:      "folder1/folder1_1/folder1_1_1/12362714-dzone-refcard215-microservices.pdf",
										Name:     "12362714-dzone-refcard215-microservices.pdf",
										Type:     "file",
										Children: nil,
									},
									"halo.txt": &feature.Object{
										Key:      "folder1/folder1_1/folder1_1_1/halo.txt",
										Name:     "halo.txt",
										Type:     "file",
										Children: nil,
									},
								},
							},
						},
					},
				},
			},
			"folder2": &feature.Object{
				Key:  "folder2/",
				Name: "folder2",
				Type: "folder",
				Children: map[string]*feature.Object{
					"BP-Diet-Roadmap-2019.pdf": &feature.Object{
						Key:      "folder2/BP-Diet-Roadmap-2019.pdf",
						Name:     "BP-Diet-Roadmap-2019.pdf",
						Type:     "file",
						Children: nil,
					},
					"folder2_1": &feature.Object{
						Key:      "folder2/folder2_1/",
						Name:     "folder2_1",
						Type:     "folder",
						Children: map[string]*feature.Object{},
					},
				},
			},
			"foto.png": {
				Key:      "foto.png",
				Name:     "foto.png",
				Type:     "file",
				Children: nil,
			},
		},
	}

	mockedS3 := s3Mock{}

	repo := &folderRepository{
		awss3: &mockedS3,
	}

	got, err := repo.ListAllObjects()

	if err != nil {
		t.Error("Unexpected error: ", err.Error())
	}

	if !reflect.DeepEqual(got, expected) {
		expectedJson, _ := json.Marshal(expected)
		gotJson, _ := json.Marshal(got)
		t.Error("\n\nExpected", string(expectedJson), "\n\nGot", string(gotJson))
	}
}

type s3MockWithError struct{}

// ListObjectsV2 returns an error
func (s *s3MockWithError) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	err := errors.New("Mock returning error")
	return &s3.ListObjectsV2Output{}, err
}
