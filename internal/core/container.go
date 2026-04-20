package core

import "context"

type ContainerPayload struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Image   string   `json:"image"`
	Command []string `json:"command"`
}
type ContainerProvider interface {
	Execute(ctx context.Context, action string, payload ContainerPayload) (string, error)
}
