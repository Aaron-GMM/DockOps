package worker

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockContainerProvider struct {
	mock.Mock
}

func (mc *MockContainerProvider) Execute(ctx context.Context, action string, payload core.ContainerPayload) (string, error) {
	args := mc.Called(ctx, action, payload)
	return args.String(0), args.Error(1)
}

func TestProcessMessage_Sucess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProvider := new(MockContainerProvider)

	fakePayload := core.ContainerPayload{
		Name:  "meu-nginx",
		Image: "nginx:latest",
	}
	msgBytes, _ := json.Marshal(fakePayload)

	mockProvider.On("Execute", mock.Anything, "create", fakePayload).
		Return("Container meu-nginx created ", nil)

	consumer := NewContainerConsumer(mockProvider)

	err := consumer.ProcessMessage(context.Background(), msgBytes)

	assert.NoError(t, err)
	mockProvider.AssertExpectations(t)
}
