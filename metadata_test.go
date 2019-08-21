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
			if !gcpmetadata.Is(err, gcpmetadata.ErrNotFoundCode) {
				t.Fatal(err)
			}
			fmt.Println("GetProjectID is NotFound")
			return
		}
	}
	fmt.Printf("GetProjectID is %v\n", p)
}

func TestExtractionRegion(t *testing.T) {
	cases := []struct {
		name       string
		metaZone   string
		wantResult string
		wantErr    gcpmetadata.ErrCode
	}{
		{"normal", "projects/999999999999/zones/asia-northeast1-1", "asia-northeast1", 0},
		{"invalid text pattern 1", "1", "", gcpmetadata.ErrInvalidArgumentCode},
		{"invalid text pattern 2", "////", "", gcpmetadata.ErrInvalidArgumentCode},
		{"empty", "", "", gcpmetadata.ErrInvalidArgumentCode},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := gcpmetadata.ExtractionRegion(tt.metaZone)
			if err != nil {
				if !gcpmetadata.Is(err, tt.wantErr) {
					t.Errorf("want error %v but got %v", tt.wantErr, err)
				}
				return
			}
			if got != tt.wantResult {
				t.Errorf("want result %v but got %v", tt.wantResult, got)
			}
		})
	}
}

func TestExtractionZone(t *testing.T) {
	cases := []struct {
		name       string
		metaZone   string
		wantResult string
		wantErr    gcpmetadata.ErrCode
	}{
		{"normal", "projects/999999999999/zones/asia-northeast1-a", "asia-northeast1-a", 0},
		{"invalid text pattern 1", "1", "1", 0},   // Zone名としてValidかがなんともいい難いので、そのまま返ってきちゃう
		{"invalid text pattern 2", "////", "", 0}, // Zone名としてValidかがなんともいい難いので、そのまま返ってきちゃう
		{"empty", "", "", gcpmetadata.ErrInvalidArgumentCode},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := gcpmetadata.ExtractionZone(tt.metaZone)
			if err != nil {
				if !gcpmetadata.Is(err, tt.wantErr) {
					t.Errorf("want error %v but got %v", tt.wantErr, err)
				}
				return
			}
			if got != tt.wantResult {
				t.Errorf("want result %v but got %v", tt.wantResult, got)
			}
		})
	}
}
