package core

import (
	"context"
	"time"
)

type EventType string

const (
	ContainerCreated EventType = "ContainerCreated"
	ContainerDeleted EventType = "ContainerDeleted"
	ContainerStarted EventType = "ContainerStarted"
	ContainerStopped EventType = "ContainerStopped"
	ContainerUpdated EventType = "ContainerUpdated"
)

// Event represents a system event related to a container or other resource.
type Event struct {
	ID         string    `json:"id" example:"evt-98765"`
	ResourceID string    `json:"resource_id" example:"cont-12345"`
	Type       EventType `json:"type" example:"ContainerCreated"`
	Payload    []byte    `json:"payload" swaggertype:"string" example:"{\"image\":\"nginx\"}"`
	CreatedAt  time.Time `json:"created_at" example:"2026-09-02T22:36:00Z"`
}

type EventRepository interface {
	Save(ctx context.Context, e Event) error
	GetByResourceID(ctx context.Context, resourceID string) ([]Event, error)
}

func DetermineContainerState(events []Event) string {
	if len(events) == 0 {
		return "Unknown"
	}
	state := "Unknown"
	for _, event := range events {
		if event.Type == ContainerCreated {
			state = "Pending"
		}
		if event.Type == ContainerDeleted {
			state = "Deleted"
		}
		if event.Type == ContainerStarted {
			state = "Running"
		}
		if event.Type == ContainerStopped {
			state = "Stopped"
		}
		if event.Type == ContainerUpdated {

		}
	}
	return state
}
