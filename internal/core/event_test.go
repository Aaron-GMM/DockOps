package core

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestDetermineContainerState(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		events         []Event
		expectedStatus string
	}{
		{
			name:           "Nenhum evento",
			events:         []Event{},
			expectedStatus: "Unknown",
		},
		{
			name: "Apenas Criado",
			events: []Event{
				{Type: ContainerCreated, CreatedAt: time.Now().Add(-2 * time.Minute)},
			},
			expectedStatus: "Pending",
		},
		{
			name: "Criado e depois Iniciado",
			events: []Event{
				{Type: ContainerCreated, CreatedAt: time.Now().Add(-2 * time.Minute)},
				{Type: ContainerStarted, CreatedAt: time.Now().Add(-1 * time.Minute)},
			},
			expectedStatus: "Running",
		},
		{
			name: "Iniciado e depois Parado",
			events: []Event{
				{Type: ContainerCreated, CreatedAt: time.Now().Add(-3 * time.Minute)},
				{Type: ContainerStarted, CreatedAt: time.Now().Add(-2 * time.Minute)},
				{Type: ContainerStopped, CreatedAt: time.Now().Add(-1 * time.Minute)},
			},
			expectedStatus: "Stopped",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := DetermineContainerState(tt.events)
			assert.Equal(t, tt.expectedStatus, status)
		})
	}
}
