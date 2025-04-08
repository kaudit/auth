package auth

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

type Authenticator interface {
	NativeAPI() (kubernetes.Interface, error)
	DynamicAPI() (dynamic.Interface, error)
}

type K8sAuthLoader interface {
	Load() ([]byte, error)
	LoadWithPath(path string) ([]byte, error)
}
