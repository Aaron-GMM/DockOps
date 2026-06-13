package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/Aaron-GMM/DockOps/internal/core"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerClient struct {
	cli *client.Client
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerClient{cli: cli}, nil
}

func (d *DockerClient) Execute(ctx context.Context, action string, payload core.ContainerPayload) (string, error) {
	switch action {
	case "create":
		return d.createAndStartContainer(ctx, payload)
	default:
		return "", fmt.Errorf("action %s not supported", action)
	}
}

func (d *DockerClient) createAndStartContainer(ctx context.Context, payload core.ContainerPayload) (string, error) {
	// 1. Pull Image
	reader, err := d.cli.ImagePull(ctx, payload.Image, types.ImagePullOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to pull image: %w", err)
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader) // Log pull progress to stdout for now

	// 2. Create Container
	resp, err := d.cli.ContainerCreate(ctx, &container.Config{
		Image: payload.Image,
		Cmd:   payload.Command,
	}, nil, nil, nil, payload.Name)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %w", err)
	}

	// 3. Start Container
	if err := d.cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container: %w", err)
	}

	return resp.ID, nil
}
