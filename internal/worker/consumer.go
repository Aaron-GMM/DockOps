package worker

import (
	"context"
	"encoding/json"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/core"
)

type ContainerConsumer struct {
	provider core.ContainerProvider
}

var log = logger.NewLogger("Consumer")

func NewContainerConsumer(provider core.ContainerProvider) *ContainerConsumer {
	return &ContainerConsumer{
		provider: provider,
	}
}

func (c *ContainerConsumer) ProcessMessage(ctx context.Context, msgBytes []byte) error {
	log.InforF("Received message of queue len: %d, message: %s", len(msgBytes), string(msgBytes))
	var payload core.ContainerPayload

	err := json.Unmarshal(msgBytes, &payload)
	if err != nil {
		log.ErrorF("Error decode json message: %w", err.Error())
		return err
	}
	log.Debugf("stared provider container payload: %+v", payload.Name)

	resultMsg, err := c.provider.Execute(ctx, "create", payload)
	if err != nil {
		log.ErrorF("Error create container: %w", err.Error())
		return err
	}
	log.InforF("Created container: %s", resultMsg)
	return nil
}
