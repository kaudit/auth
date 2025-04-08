# Kubernetes Authentication Library

A Go library for standardized Kubernetes client authentication and configuration management.

## Overview

This library provides a clean, modular approach to handling Kubernetes authentication. It abstracts away the details of loading kubeconfig files and establishing connections to Kubernetes clusters, making it easier to build applications that interact with Kubernetes.

## Features

- Abstract interfaces for Kubernetes authentication
- File-based kubeconfig loading
- Support for both typed and dynamic Kubernetes clients
- Validation of configuration files

## Installation

Make sure you have Go 1.23 or later installed.

To add this library to your Go project, use the `go get` command:

```bash
# Install dependencies
go get github.com/kaudit/val
go get k8s.io/client-go
go get k8s.io/api
go get k8s.io/apimachinery

# Install the main library
go get github.com/kaudit/auth
```

## Usage

### Basic Usage

```go
// Create a kubeconfig loader
loader := k8s_auth_data_loader.NewK8sConfigLoader("/path/to/kubeconfig")

// Create an authenticator
authenticator, err := kube_config.NewKubeConfigAuthenticator(loader)
if err != nil {
    // Handle error
}

// Get a typed Kubernetes client
clientset, err := authenticator.NativeAPI()
if err != nil {
    // Handle error
}

// Use the client
pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})

// Or get a dynamic client
dynamicClient, err := authenticator.DynamicAPI()
if err != nil {
    // Handle error
}
```

### Custom Configuration Path

You can also load a kubeconfig from a specific path:

```go
// Create a kubeconfig loader with default path
loader := k8s_auth_data_loader.NewK8sConfigLoader("~/.kube/config")

// Load a specific config file
configData, err := loader.LoadWithPath("/path/to/another/kubeconfig")
if err != nil {
    // Handle error
}
```
## API Documentation

### Authenticator

#### `NativeAPI() (kubernetes.Interface, error)`
Returns a typed Kubernetes client constructed from authenticated configuration.
- Returns a fully configured Kubernetes client interface or an error if client creation fails

#### `DynamicAPI() (dynamic.Interface, error)`
Returns a dynamic Kubernetes client constructed from authenticated configuration.
- Returns a fully configured dynamic client interface or an error if client creation fails

### K8sAuthLoader

#### `Load() ([]byte, error)`
Reads and returns kubeconfig data from the loader's configured default path.
- Returns raw kubeconfig data as a byte array or an error if loading fails

#### `LoadWithPath(path string) ([]byte, error)`
Reads and returns kubeconfig data from the specified path.
- `path`: File path to a kubeconfig file (must be valid and accessible)
- Returns raw kubeconfig data as a byte array or an error if loading fails

### KubeConfig

#### `NewKubeConfigAuthenticator(loader auth.K8sAuthLoader) (auth.Authenticator, error)`
Returns an implementation of the auth.Authenticator interface.
- `loader`: A K8sAuthLoader implementation that will provide kubeconfig data
- Returns a configured authenticator or an error if initialization fails

#### `NativeAPI() (kubernetes.Interface, error)`
Returns a typed Kubernetes client constructed from kubeconfig data.
- Uses the configured auth loader to obtain configuration
- Returns a fully configured Kubernetes client interface or an error if creation fails

#### `DynamicAPI() (dynamic.Interface, error)`
Returns a dynamic Kubernetes client constructed from kubeconfig data.
- Uses the configured auth loader to obtain configuration
- Returns a fully configured dynamic client interface or an error if creation fails

### K8sAuthDataLoader

#### `NewK8sConfigLoader(path string) auth.K8sAuthLoader`
Returns a new instance of K8sAuthDataLoader.
- `path`: Default file path to use for kubeconfig loading
- Returns a configured K8sAuthLoader implementation

#### `Load() ([]byte, error)`
Reads and returns kubeconfig data from the loader's configured path.
- Validates that the path refers to an accessible file
- Returns raw kubeconfig data as a byte array or an error if validation or reading fails

#### `LoadWithPath(path string) ([]byte, error)`
Reads and returns kubeconfig data from the specified path.
- `path`: File path to a kubeconfig file (must be valid and accessible)
- Validates that the path refers to an accessible file
- Returns raw kubeconfig data as a byte array or an error if validation or reading fails

## Architecture

The library is built around several key components:

- **Authenticator**: Interface for accessing Kubernetes API clients
- **K8sAuthLoader**: Interface for loading Kubernetes authentication data
- **KubeConfig**: Implementation of the Authenticator interface using config files
- **K8sAuthDataLoader**: Implementation of the K8sAuthLoader interface for file-based loading

## License

This project is licensed under the [MIT License](./LICENSE)

## Thanks

We would like to express our gratitude to the Kubernetes team and contributors for creating and maintaining the excellent `k8s.io` packages that this library builds upon. Their work on `client-go`, `api`, and `apimachinery` provides the robust foundation that makes this authentication abstraction layer possible.

