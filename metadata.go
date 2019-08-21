package gcpmetadata

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"github.com/morikuni/failure"
)

// OnGCP is GCP上で動いているかどうかを返す
// GCP上と判断されるか確認したのは以下
// Google App Engine Standard for Go 1.11
// Google Compute Engine
// Google Kubernetes Engine
func OnGCP() bool {
	return metadata.OnGCE()
}

// GetProjectID is Return current GCP ProjectID
// GCP上で動いている場合は、Project Metadataから取得し、そうでなければ、環境変数から取得する
func GetProjectID() (string, error) {
	if !metadata.OnGCE() {
		p := os.Getenv("GOOGLE_CLOUD_PROJECT")
		if p != "" {
			return p, nil
		}
		p = os.Getenv("GCLOUD_PROJECT")
		if p != "" {
			return p, nil
		}
		return "", ErrNotFound("project id environment valiable is not found. plz set $GOOGLE_CLOUD_PROJECT")
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return "", err
	}
	if projectID == "" {
		return "", ErrNotFound("project id is not found")
	}
	return projectID, nil
}

// GetServiceAccountEmail is Return current Service Account Email
// GCP上で動いている場合は、Metadataから取得し、そうでなければ、環境変数から取得する
func GetServiceAccountEmail() (string, error) {
	if !metadata.OnGCE() {
		return os.Getenv("GCLOUD_SERVICE_ACCOUNT"), nil
	}
	sa, err := getMetadata("service-accounts/default/email")
	if err != nil {
		return "", failure.Wrap(err, failure.Message("failed get ServiceAccountEmail"))
	}
	return string(sa), nil
}

// GetRegion is Appが動いているRegionを取得する
func GetRegion() (string, error) {
	if !metadata.OnGCE() {
		return os.Getenv("INSTANCE_REGION"), nil
	}
	zone, err := getMetadata("zone")
	if err != nil {
		return "", failure.Wrap(err, failure.Message("failed get Zone"))
	}

	return ExtractionRegion(string(zone))
}

// GetZone is Appが動いているZoneを取得する
func GetZone() (string, error) {
	if !metadata.OnGCE() {
		return os.Getenv("INSTANCE_ZONE"), nil
	}
	zone, err := getMetadata("zone")
	if err != nil {
		return "", failure.Wrap(err, failure.Message("failed get Zone"))
	}

	return ExtractionZone(string(zone))
}

// ExtractionRegion is Metadata Serverから取得する projects/[NUMERIC_PROJECT_ID]/zones/[ZONE] 形式の文字列から、Region部分を取り出す
func ExtractionRegion(metaZone string) (string, error) {
	l := strings.Split(string(metaZone), "/")
	if len(l) < 1 {
		return "", ErrInvalidArgument("projects/[NUMERIC_PROJECT_ID]/zones/[ZONE]", metaZone)
	}
	v := l[len(l)-1]
	if len(v) < 3 {
		return "", ErrInvalidArgument("projects/[NUMERIC_PROJECT_ID]/zones/[ZONE]", metaZone)
	}
	v = v[:len(v)-2]
	return v, nil
}

// ExtractionZone is Metadata Serverから取得する projects/[NUMERIC_PROJECT_ID]/zones/[ZONE] 形式の文字列から、Zone部分を取り出す
func ExtractionZone(metaZone string) (string, error) {
	l := strings.Split(string(metaZone), "/")
	if len(l) < 1 {
		return "", ErrInvalidArgument("projects/[NUMERIC_PROJECT_ID]/zones/[ZONE]", metaZone)
	}
	return l[len(l)-1], nil
}

// GetInstanceAttribute is Instance Metadataを取得する
// GCP以外で動いている時は、環境変数を取得する
func GetInstanceAttribute(key string) (string, error) {
	if !metadata.OnGCE() {
		return os.Getenv(fmt.Sprintf("INSTANCE_%s", key)), nil
	}

	v, err := metadata.InstanceAttributeValue(key)
	if err != nil {
		return "", err
	}
	return v, nil
}

// GetProjectAttribute is Project Metadataを取得する
// GCP以外で動いている時は、環境変数を取得する
func GetProjectAttribute(key string) (string, error) {
	if !metadata.OnGCE() {
		return os.Getenv(fmt.Sprintf("PROJECT_%s", key)), nil
	}

	v, err := metadata.ProjectAttributeValue(key)
	if err != nil {
		return "", err
	}
	return v, nil
}

func getMetadata(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://metadata.google.internal/computeMetadata/v1/instance/%s", path), nil)
	if err != nil {
		return nil, failure.Wrap(err, failure.Messagef("failed http.NewRequest. path=%s", path))
	}
	req.Header.Set("Metadata-Flavor", "Google")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, failure.Wrap(err, failure.Messagef("failed http.SendReq. path=%s", path))
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, failure.Wrap(err, failure.Messagef("failed read response.Body. path=%s", path))
	}
	if res.StatusCode != http.StatusOK {
		return nil, failure.Unexpected(fmt.Sprintf("metadata server response is %v:%v", res.StatusCode, string(b)))
	}

	return b, nil
}
