package core

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestDetermineContainerState_SemEventos_DeveRetornarUnknown(t *testing.T) {
	// Arrange
	eventos := []Event{}

	// Act
	status := DetermineContainerState(eventos)

	// Assert
	assert.Equal(t, "Unknown", status)
}
func TestDetermineContainerState_ApenasEventoCriado_DeveRetornarPending(t *testing.T) {
	// Arrange
	eventos := []Event{
		{Type: ContainerCreated, CreatedAt: time.Now().Add(-2 * time.Minute)},
	}

	// Act
	status := DetermineContainerState(eventos)

	// Assert
	assert.Equal(t, "Pending", status)
}
func TestDetermineContainerState_EventosCriadoEIniciado_DeveRetornarRunning(t *testing.T) {
	// Arrange
	eventos := []Event{
		{Type: ContainerCreated, CreatedAt: time.Now().Add(-2 * time.Minute)},
		{Type: ContainerStarted, CreatedAt: time.Now().Add(-1 * time.Minute)},
	}

	// Act
	status := DetermineContainerState(eventos)

	// Assert
	assert.Equal(t, "Running", status)
}
func TestDetermineContainerState_EventosAteParado_DeveRetornarStopped(t *testing.T) {
	// Arrange
	eventos := []Event{
		{Type: ContainerCreated, CreatedAt: time.Now().Add(-3 * time.Minute)},
		{Type: ContainerStarted, CreatedAt: time.Now().Add(-2 * time.Minute)},
		{Type: ContainerStopped, CreatedAt: time.Now().Add(-1 * time.Minute)},
	}

	// Act
	status := DetermineContainerState(eventos)

	// Assert
	assert.Equal(t, "Stopped", status)
}
func TestDetermineContainerState_EventosAteDeletado_DeveRetornarDeleted(t *testing.T) {
	// Arrange
	eventos := []Event{
		{Type: ContainerCreated, CreatedAt: time.Now().Add(-3 * time.Minute)},
		{Type: ContainerStarted, CreatedAt: time.Now().Add(-2 * time.Minute)},
		{Type: ContainerDeleted, CreatedAt: time.Now().Add(-1 * time.Minute)},
	}

	// Act
	status := DetermineContainerState(eventos)

	// Assert
	assert.Equal(t, "Deleted", status)
}
