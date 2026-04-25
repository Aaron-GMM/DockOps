package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

type MockPublisher struct {
	mock.Mock
}

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) Save(ctx context.Context, event core.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *MockPublisher) Publish(ctx context.Context, queueName string, message []byte) error {
	args := m.Called(ctx, queueName, message)
	return args.Error(0)
}

func TestCreateContainer_ErroAoFazerBindDoJSON_DeveRetornarBadRequest(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockPub := new(MockPublisher)
	mockRep := new(MockEventRepository)

	// Nenhum mock é configurado para ser chamado, pois deve falhar antes

	handler := NewContainerHandler(mockPub, mockRep)
	router := gin.Default()
	router.POST("/api/v1/containers", handler.CreateContainer)

	jsonInvalido := []byte(`{"name": "meu-nginx", "image": }`) // JSON quebrado
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/containers", bytes.NewBuffer(jsonInvalido))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
	mockRep.AssertNotCalled(t, "Save")
	mockPub.AssertNotCalled(t, "Publish")
}
func TestCreateContainer_ErroAoPublicarNaFila_DeveRetornarInternalServerError(t *testing.T) {
	// Arrange
	gin.SetMode(gin.TestMode)
	mockPub := new(MockPublisher)
	mockRep := new(MockEventRepository)

	mockRep.On("Save", mock.Anything, mock.AnythingOfType("core.Event")).Return(nil)
	// Simulamos um erro na mensageria
	mockPub.On("Publish", mock.Anything, "container_tasks", mock.Anything).Return(errors.New("rabbitmq offline"))

	handler := NewContainerHandler(mockPub, mockRep)
	router := gin.Default()
	router.POST("/api/v1/containers", handler.CreateContainer)

	jsonPayload := []byte(`{"name": "meu-nginx", "image": "nginx:latest"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/containers", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Act
	router.ServeHTTP(w, req)

	// Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockRep.AssertExpectations(t)
	mockPub.AssertExpectations(t)
}
func TestCreateContainer_ErrorAoSalvarNoBanco_DeveRetornarInternalServerError(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)
	mockPublisher := new(MockPublisher)
	mockRepository := new(MockEventRepository)

	mockRepository.On("Save", mock.Anything, mock.AnythingOfType("core.Event")).Return(errors.New("data base connection lost"))

	handler := NewContainerHandler(mockPublisher, mockRepository)

	router := gin.Default()
	router.POST("/api/v1/container", handler.CreateContainer)
	jsonPayload := []byte(`{"name": "meu-nginx", "image":"nginx:latest"}`)
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/container", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	//Act
	router.ServeHTTP(w, req)

	//Assert
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mockPublisher.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}
func TestCreateContainer_Successo_DeveRetornarStatusAccepted(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	mockPub := new(MockPublisher)
	mockRep := new(MockEventRepository)

	mockRep.On("Save", mock.Anything, mock.AnythingOfType("core.Event")).Return(nil)
	mockPub.On("Publish", mock.Anything, "container_tasks", mock.Anything).Return(nil)

	containerHandler := NewContainerHandler(mockPub, mockRep)

	router := gin.Default()
	router.POST("api/v1/containers", containerHandler.CreateContainer)

	jsonPayload := []byte(`{"name": "meu-nginx", "image": "nginx:latest"}`)
	req, _ := http.NewRequest("POST", "/api/v1/containers", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	//Act
	router.ServeHTTP(w, req)
	//Assert
	assert.Equal(t, http.StatusAccepted, w.Code)
	mockPub.AssertExpectations(t)
	mockRep.AssertExpectations(t)
}
