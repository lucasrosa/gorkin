package s3

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

func TestListChildren(t *testing.T) {

	prefix := "folder1"
	key := "object1"

	outputObject := s3.ListObjectsV2Output{
		CommonPrefixes: []*s3.CommonPrefix{
			&s3.CommonPrefix{
				Prefix: &prefix,
			},
		},
		Contents: []*s3.Object{
			&s3.Object{
				Key: &key,
			},
		},
	}

	expected := []string{"folder1", "object1"}

	got := listChildren(&outputObject, "folder0")

	if !reflect.DeepEqual(got, expected) {
		t.Error("Expected", expected, "got", got)
	}

}
