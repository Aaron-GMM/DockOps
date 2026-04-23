package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

func (m *MockPublisher) Publish(ctx context.Context, queueName string, message []byte) error {
	args := m.Called(ctx, queueName, message)
	return args.Error(0)
}

func TestCreateContainer_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockPub := new(MockPublisher)

	mockPub.On("Publish", mock.Anything, "container_tasks", mock.Anything).Return(nil)

	containerHandler := NewContainerHandler(mockPub)

	router := gin.Default()
	router.POST("api/v1/containers", containerHandler.CreateContainer)

	jsonPayload := []byte(`{"name": "meu-nginx", "image": "nginx:latest"}`)
	req, _ := http.NewRequest("POST", "/api/v1/containers", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusAccepted, w.Code)
	mockPub.AssertExpectations(t)
}
