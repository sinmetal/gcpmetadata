package gcpmetadata

import "os"

const AppEngineService = "GAE_SERVICE"
const AppEngineVersion = "GAE_VERSION"
const AppEngineInstance = "GAE_INSTANCE"
const AppEngineRuntime = "GAE_RUNTIME"
const AppEngineMemoryMB = "GAE_MEMORY_MB"
const AppEngineDeploymentID = "GAE_DEPLOYMENT_ID"
const AppEngineEnv = "GAE_ENV"

// GetAppEngineService is return service id
// The service name specified in your app.yaml file. If no service name is specified, it is set to default.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineService() (string, error) {
	v := os.Getenv(AppEngineService)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Service id environment valiable is not found. plz set $GAE_SERVICE")
}

// GetAppEngineVersion is return version id
// The current version label of your service.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineVersion() (string, error) {
	v := os.Getenv(AppEngineVersion)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Version id environment valiable is not found. plz set $GAE_VERSION")
}

// GetAppEngineInstance is return version id
// The ID of the instance on which your service is currently running.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineInstance() (string, error) {
	v := os.Getenv(AppEngineInstance)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Instance id environment valiable is not found. plz set $GAE_INSTANCE")
}

// GetAppEngineRuntime is return runtime
// The runtime specified in your app.yaml file.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineRuntime() (string, error) {
	v := os.Getenv(AppEngineRuntime)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Runtime id environment valiable is not found. plz set $GAE_RUNTIME")
}

// GetAppEngineMemoryMB is return MemoryMB
// The amount of memory available to the application process, in MB.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineMemoryMB() (string, error) {
	v := os.Getenv(AppEngineMemoryMB)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine MemoryMB id environment valiable is not found. plz set $GAE_MEMORY_MB")
}

// GetAppEngineDeploymentID is return deployment id
// The ID of the current deployment.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineDeploymentID() (string, error) {
	v := os.Getenv(AppEngineDeploymentID)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Deployment id environment valiable is not found. plz set $GAE_DEPLOYMENT_ID")
}

// GetAppEngineEnv is return env
// The App Engine environment. Set to standard.
// https://cloud.google.com/appengine/docs/standard/go/runtime#environment_variables
func GetAppEngineEnv() (string, error) {
	v := os.Getenv(AppEngineEnv)
	if v != "" {
		return v, nil
	}
	return "", ErrNotFound("AppEngine Deployment id environment valiable is not found. plz set $GAE_ENV")
}
