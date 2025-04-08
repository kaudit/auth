package k8s_auth_data_loader

import (
	"fmt"
	"os"

	"github.com/kaudit/val"

	"github.com/kaudit/auth"
)

// K8sAuthDataLoader implements the auth.K8sAuthLoader interface.
// It loads Kubernetes authentication data from a specified file path on the filesystem.
type K8sAuthDataLoader struct {
	path string
}

// NewK8sConfigLoader returns a new instance of K8sAuthDataLoader.
// It uses the provided path as the default source for kubeconfig data.
func NewK8sConfigLoader(path string) auth.K8sAuthLoader {
	return &K8sAuthDataLoader{path: path}
}

// Load reads and returns kubeconfig data from the loader's configured path.
// It returns an error if validation or file reading fails.
func (k *K8sAuthDataLoader) Load() ([]byte, error) {
	err := val.ValidateWithTag(k.path, "required,file")
	if err != nil {
		return nil, fmt.Errorf("val.ValidateWithTag failed: %w", err)
	}

	b, err := os.ReadFile(k.path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile failed: %w", err)
	}

	return b, nil
}

// LoadWithPath reads and returns kubeconfig data from the specified path.
// It returns an error if validation or file reading fails.
func (k *K8sAuthDataLoader) LoadWithPath(path string) ([]byte, error) {
	err := val.ValidateWithTag(path, "required,file")
	if err != nil {
		return nil, fmt.Errorf("val.ValidateWithTag failed: %w", err)
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile failed: %w", err)
	}

	return b, nil
}
