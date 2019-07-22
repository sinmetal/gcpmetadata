package gcpmetadata_test

import (
	"fmt"
	"testing"

	"github.com/sinmetal/gcpmetadata"
)

func TestOnGCP(t *testing.T) {
	v := gcpmetadata.OnGCP()
	fmt.Printf("OnGCP is %v\n", v)
}

func TestGetProjectID(t *testing.T) {
	onGCP := gcpmetadata.OnGCP()

	p, err := gcpmetadata.GetProjectID()
	if onGCP {
		if err != nil {
			t.Fatal(err)
		}
	} else {
		if err != nil {
			if err != gcpmetadata.ErrNotFound {
				t.Fatal(err)
			}
			fmt.Println("GetProjectID is NotFound")
			return
		}
	}
	fmt.Printf("GetProjectID is %v\n", p)
}
