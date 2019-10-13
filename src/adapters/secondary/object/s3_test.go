package s3

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
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

type s3MockWithError struct{}

// ListObjectsV2 returns an error
func (s *s3MockWithError) ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	err := errors.New("Mock returning error")
	return &s3.ListObjectsV2Output{}, err
}
