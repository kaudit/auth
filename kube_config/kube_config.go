package kube_config

import (
	"fmt"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kaudit/auth"
)

type KubeConfig struct {
	authLoader auth.K8sAuthLoader
}

func NewKubeConfigAuthenticator(loader auth.K8sAuthLoader) (auth.Authenticator, error) {
	return &KubeConfig{authLoader: loader}, nil
}

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
