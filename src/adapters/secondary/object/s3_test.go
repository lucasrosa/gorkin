package s3

import (
	"errors"
	"reflect"
	"testing"

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

// ListObjectsV2 returns an object with 2 folders and 2 files
func (s *s3Mock) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	prefix := "folder1"
	prefix2 := "folder2"
	key := "object1"
	key2 := "object2"

	outputObject := s3.ListObjectsV2Output{
		CommonPrefixes: []*s3.CommonPrefix{
			&s3.CommonPrefix{
				Prefix: &prefix,
			},
			&s3.CommonPrefix{
				Prefix: &prefix2,
			},
		},
		Contents: []*s3.Object{
			&s3.Object{
				Key: &key,
			},
			&s3.Object{
				Key: &key2,
			},
		},
	}

	return &outputObject, nil
}

func TestListObjects(t *testing.T) {
	mockedS3 := s3Mock{}
	repo := &folderRepository{&mockedS3}

	folder := "object2"

	got, err := repo.ListObjects(folder)

	if err != nil {
		t.Error("No error expected. Got", err)
	}

	children := []string{"folder1", "folder2", "object1"}
	expected := feature.Folder{
		Children: children,
	}

	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected", expected, "got", got)
	}
}

type s3MockWithError struct{}

// ListObjectsV2 returns an error
func (s *s3MockWithError) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	err := errors.New("Mock returning error")
	return &s3.ListObjectsV2Output{}, err
}

func TestListObjectsWithError(t *testing.T) {
	mockedS3 := s3MockWithError{}
	repo := &folderRepository{&mockedS3}

	folder := "object2"

	_, err := repo.ListObjects(folder)

	if err == nil {
		t.Error("No error expected. Got", err)
	}

	expected := errors.New("Mock returning error")

	if err.Error() != expected.Error() {
		t.Error("Expected", expected, "got", err)
	}
}
