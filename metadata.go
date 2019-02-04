package gcpmetadata

import (
	"errors"
	"fmt"
	"os"

	"cloud.google.com/go/compute/metadata"
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
		return os.Getenv("GCLOUD_PROJECT"), nil
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
