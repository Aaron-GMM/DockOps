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

type Event struct {
	ID         string    `json:"id"`
	ResourceID string    `json:"resource_id"`
	Type       EventType `json:"type"`
	Payload    []byte    `json:"payload"`
	CreatedAt  time.Time `json:"created_at"`
}

type EventRepository interface {
	Save(ctx context.Context, e Event) error
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
