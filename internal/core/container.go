package core

import "context"

// ContainerPayload represents the data needed to create or manage a container.
type ContainerPayload struct {
	ID      string            `json:"id" example:"cont-12345"`
	Name    string            `json:"name" example:"my-nginx"`
	Image   string            `json:"image" example:"nginx:latest"`
	Command []string          `json:"command" example:"[\"nginx\", \"-g\", \"daemon off;\"]"`
	Env     []string          `json:"env" example:"[\"FOO=BAR\",\"DB_PASS=123\"]"`
	Ports   map[string]string `json:"ports" example:"{\"80/tcp\":\"8080\"}"`
}
type ContainerProvider interface {
	Execute(ctx context.Context, action string, payload ContainerPayload) (string, error)
}
