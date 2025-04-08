package kubeconfig

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	mocksauth "github.com/kaudit/auth/mocks/K8sAuthLoader"
)

func TestNewKubeConfigAuthenticator(t *testing.T) {
	// Arrange
	mockLoader := &mocksauth.MockK8sAuthLoader{}

	// Act
	authenticator, err := NewKubeConfigAuthenticator(mockLoader)

	// Assert
	require.NoError(t, err)
	assert.NotNil(t, authenticator)
}

func TestKubeConfig_NativeAPI(t *testing.T) {
	t.Run("successful client creation", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		validKubeconfig := []byte(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://kubernetes.default.svc
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
users:
- name: test-user
  user: {}
`)

		mockLoader.EXPECT().Load().Return(validKubeconfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.NativeAPI()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, client)
		assert.Implements(t, (*kubernetes.Interface)(nil), client)
	})

	t.Run("loader error", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		mockLoader.EXPECT().Load().Return(nil, errors.New("load error"))

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.NativeAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "authLoader.Load failed")
		assert.Nil(t, client)
	})

	t.Run("invalid client config", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		// Valid YAML but missing current-context which will cause clientCfg.ClientConfig() to fail
		invalidClientConfig := []byte(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://kubernetes.default.svc
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
# current-context is missing
users:
- name: test-user
  user: {}
`)

		mockLoader.EXPECT().Load().Return(invalidClientConfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.NativeAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "clientCfg.ClientConfig failed")
		assert.Nil(t, client)
	})

	t.Run("invalid kubeconfig", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		invalidKubeconfig := []byte(`invalid yaml`)

		mockLoader.EXPECT().Load().Return(invalidKubeconfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.NativeAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "getRestConfig failed")
		assert.Nil(t, client)
	})
}

func TestKubeConfig_DynamicAPI(t *testing.T) {
	t.Run("successful client creation", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		validKubeconfig := []byte(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://kubernetes.default.svc
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
users:
- name: test-user
  user: {}
`)

		mockLoader.EXPECT().Load().Return(validKubeconfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.DynamicAPI()

		// Assert
		require.NoError(t, err)
		assert.NotNil(t, client)
		assert.Implements(t, (*dynamic.Interface)(nil), client)
	})

	t.Run("loader error", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		mockLoader.EXPECT().Load().Return(nil, errors.New("load error"))

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.DynamicAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "authLoader.Load failed")
		assert.Nil(t, client)
	})

	t.Run("invalid client config", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		// Valid YAML but missing current-context which will cause clientCfg.ClientConfig() to fail
		invalidClientConfig := []byte(`
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: https://kubernetes.default.svc
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
# current-context is missing
users:
- name: test-user
  user: {}
`)

		mockLoader.EXPECT().Load().Return(invalidClientConfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.DynamicAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "clientCfg.ClientConfig failed")
		assert.Nil(t, client)
	})

	t.Run("invalid kubeconfig", func(t *testing.T) {
		// Arrange
		mockLoader := &mocksauth.MockK8sAuthLoader{}
		invalidKubeconfig := []byte(`invalid yaml`)

		mockLoader.EXPECT().Load().Return(invalidKubeconfig, nil)

		authenticator, err := NewKubeConfigAuthenticator(mockLoader)
		require.NoError(t, err)

		// Act
		client, err := authenticator.DynamicAPI()

		// Assert
		require.Error(t, err)
		assert.Contains(t, err.Error(), "getRestConfig failed")
		assert.Nil(t, client)
	})
}
