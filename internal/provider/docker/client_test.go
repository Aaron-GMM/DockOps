package docker

import (
	"context"
	"testing"

	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDockerClient is a simplified mock for the Docker SDK client
// In a real scenario, we might want to mock the specific methods we use.
type MockDockerClient struct {
	mock.Mock
}

// For simplicity in this environment, I'll mock the Execute method of our DockerClient wrapper 
// if I can't easily mock the underlying SDK without many interfaces.
// But the goal is to test DockerClient, so I should mock the SDK.

func TestDockerClient_Execute_UnsupportedAction(t *testing.T) {
	// Arrange
	d := &DockerClient{}
	payload := core.ContainerPayload{}

	// Act
	_, err := d.Execute(context.Background(), "invalid", payload)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not supported")
}

// Since mocking the Docker SDK client requires implementing many methods (it's a huge interface),
// and I don't have a mock generator here, I'll focus on the logic I can test.
// A better way would be to wrap the SDK client in an interface we control.
