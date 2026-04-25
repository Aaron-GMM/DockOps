package worker

import (
	"context"
	"encoding/json"
	"errors"
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

type MockEventRepository struct {
	mock.Mock
}

func (me *MockEventRepository) Save(ctx context.Context, event core.Event) error {
	args := me.Called(ctx, event)
	return args.Error(0)
}

func TestProcessMessage_ErroAoDecodificarJSON_DeveRetornarErro(t *testing.T) {
	// Arrange
	mockProvider := new(MockContainerProvider)
	mockRepository := new(MockEventRepository)
	consumer := NewContainerConsumer(mockProvider, mockRepository)
	msgBytesQuebrada := []byte(`{"id": "123", "name": }`) // JSON inválido

	// Act
	err := consumer.ProcessMessage(context.Background(), msgBytesQuebrada)

	// Assert
	assert.Error(t, err)
	mockProvider.AssertNotCalled(t, "Execute")
	mockRepository.AssertNotCalled(t, "Save")
}
func TestProcessMessage_ErroNoProvider_DeveRetornarErro(t *testing.T) {
	// Arrange
	mockProvider := new(MockContainerProvider)
	mockRepository := new(MockEventRepository)

	fakePayload := core.ContainerPayload{
		ID:    "123",
		Name:  "meu-nginx",
		Image: "nginx:latest",
	}
	msgBytes, _ := json.Marshal(fakePayload)

	// Simulamos uma falha ao tentar criar o container no Docker
	mockProvider.On("Execute", mock.Anything, "create", fakePayload).Return("", errors.New("docker daemon not running"))

	consumer := NewContainerConsumer(mockProvider, mockRepository)

	// Act
	err := consumer.ProcessMessage(context.Background(), msgBytes)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "docker daemon not running")
	mockProvider.AssertExpectations(t)
	mockRepository.AssertNotCalled(t, "Save") // Se o container falhou ao criar, não salvamos ContainerStarted
}
func TestProcessMessage_ErroAoSalvarEvento_DeveRetornarErro(t *testing.T) {
	// Arrange
	mockProvider := new(MockContainerProvider)
	mockRepository := new(MockEventRepository)

	fakePayload := core.ContainerPayload{
		ID:    "123",
		Name:  "meu-nginx",
		Image: "nginx:latest",
	}
	msgBytes, _ := json.Marshal(fakePayload)

	mockProvider.On("Execute", mock.Anything, "create", fakePayload).Return("Created", nil)
	// Simulamos falha no banco de dados ao salvar o evento
	mockRepository.On("Save", mock.Anything, mock.AnythingOfType("core.Event")).Return(errors.New("db error"))

	consumer := NewContainerConsumer(mockProvider, mockRepository)

	// Act
	err := consumer.ProcessMessage(context.Background(), msgBytes)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "db error")
	mockProvider.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}
func TestProcessMessage_Sucesso_DeveRetornarNil(t *testing.T) {
	//Arrange
	gin.SetMode(gin.TestMode)

	mockProvider := new(MockContainerProvider)
	mockRepository := new(MockEventRepository)

	fakePayload := core.ContainerPayload{
		ID:    "meu-id_simulado-123",
		Name:  "meu-nginx",
		Image: "nginx:latest",
	}
	msgBytes, _ := json.Marshal(fakePayload)

	mockProvider.On("Execute", mock.Anything, "create", fakePayload).
		Return("Container meu-nginx created ", nil)
	mockRepository.On("Save", mock.Anything, mock.MatchedBy(func(e core.Event) bool {
		return e.Type == core.ContainerStarted && e.ResourceID == "meu-id_simulado-123"
	})).Return(nil)

	consumer := NewContainerConsumer(mockProvider, mockRepository)
	//Act
	err := consumer.ProcessMessage(context.Background(), msgBytes)

	//Assert
	assert.NoError(t, err)
	mockProvider.AssertExpectations(t)
	mockRepository.AssertExpectations(t)
}
