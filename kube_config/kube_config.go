package kubeconfig

import (
	"fmt"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kaudit/auth"
)

// KubeConfig implements the auth.Authenticator interface using a K8sAuthLoader.
// It loads kubeconfig data on demand and constructs both typed and dynamic clients from it.
type KubeConfig struct {
	authLoader auth.K8sAuthLoader
}

// NewKubeConfigAuthenticator returns an implementation of the auth.Authenticator interface.
// It uses the provided K8sAuthLoader to load kubeconfig data on demand.
func NewKubeConfigAuthenticator(loader auth.K8sAuthLoader) (auth.Authenticator, error) {
	return &KubeConfig{authLoader: loader}, nil
}

// NativeAPI returns a typed Kubernetes client constructed from kubeconfig data.
// It returns an error if loading the configuration or creating the client fails.
func (k *KubeConfig) NativeAPI() (kubernetes.Interface, error) {
	kubeConfig, err := k.authLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("authLoader.Load failed: %w", err)
	}

	r, err := getRestConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("getRestConfig failed: %w", err)
	}

	i, err := kubernetes.NewForConfig(r)
	if err != nil {
		return nil, fmt.Errorf("kubernetes.NewForConfig failed: %w", err)
	}

	return i, nil
}

// DynamicAPI returns a dynamic Kubernetes client constructed from kubeconfig data.
// It returns an error if loading the configuration or creating the client fails.
func (k *KubeConfig) DynamicAPI() (dynamic.Interface, error) {
	kubeConfig, err := k.authLoader.Load()
	if err != nil {
		return nil, fmt.Errorf("authLoader.Load failed: %w", err)
	}

	r, err := getRestConfig(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("getRestConfig failed: %w", err)
	}

	d, err := dynamic.NewForConfig(r)
	if err != nil {
		return nil, fmt.Errorf("dynamic.NewForConfig failed: %w", err)
	}

	return d, nil
}

// getRestConfig constructs a *rest.Config object from the given kubeconfig data.
// It returns an error if the kubeconfig is invalid or cannot be parsed.
func getRestConfig(kubeConfig []byte) (*rest.Config, error) {
	cfg, err := clientcmd.Load(kubeConfig)
	if err != nil {
		return nil, fmt.Errorf("clientcmd.Load failed: %w", err)
	}

	clientCfg := clientcmd.NewDefaultClientConfig(*cfg, &clientcmd.ConfigOverrides{})

	restCfg, err := clientCfg.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("clientCfg.ClientConfig failed: %w", err)
	}

	return restCfg, nil
}
