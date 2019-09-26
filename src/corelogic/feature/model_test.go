package feature

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {

	a := Object{
		ID:   "1",
		Name: "dasd a",
		Type: "folder",
		Children: []Object{
			Object{
				ID:   "2",
				Name: "dassdasd",
				Type: "file",
			},
			Object{
				ID:   "3",
				Name: "fdsfsd32",
				Type: "folder",
				Children: []Object{
					Object{
						ID:   "4",
						Name: "dassdasd",
						Type: "file",
					},
				},
			},
		},
	}
	b, err := json.Marshal(a)

	if err != nil {
		t.Error("fedeu")
	}

	fmt.Printf("%s\n", b)
	expected := 1

	if 2 != expected {
		t.Error("Expected", expected, "didnt get")
	}
}
