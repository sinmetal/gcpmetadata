package gcpmetadata

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/morikuni/failure"
)

var ErrNotFound = errors.New("not found")

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
		return "", ErrNotFound
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		return "", err
	}
	if projectID == "" {
		return "", ErrNotFound
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
