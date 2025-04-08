package k8s_auth_data_loader

import (
	"fmt"
	"os"

	"github.com/kaudit/val"

	"github.com/kaudit/auth"
)

type K8sAuthDataLoader struct {
	path string
}

func NewK8sConfigLoader(path string) auth.K8sAuthLoader {
	return &K8sAuthDataLoader{path: path}
}

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
