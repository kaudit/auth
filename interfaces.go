package auth

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// Authenticator defines an abstraction for accessing Kubernetes API clients.
// It provides methods to construct and retrieve both typed and dynamic clients
// based on a preconfigured authentication source.
type Authenticator interface {
	NativeAPI() (kubernetes.Interface, error)
	DynamicAPI() (dynamic.Interface, error)
}

// K8sAuthLoader defines a mechanism for loading Kubernetes authentication configuration data.
// It supports loading from both a default source and a user-specified file path.
type K8sAuthLoader interface {
	Load() ([]byte, error)
	LoadWithPath(path string) ([]byte, error)
}
