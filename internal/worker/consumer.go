package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	"github.com/Aaron-GMM/DockOps/internal/core"
)

type ContainerConsumer struct {
	provider   core.ContainerProvider
	repository core.EventRepository
}

var log = logger.NewLogger("Consumer")

func NewContainerConsumer(provider core.ContainerProvider, repository core.EventRepository) *ContainerConsumer {
	return &ContainerConsumer{
		provider:   provider,
		repository: repository,
	}
}

func (c *ContainerConsumer) ProcessMessage(ctx context.Context, msgBytes []byte) error {
	log.InforF("Received message of queue len: %d, message: %s", len(msgBytes), string(msgBytes))
	var payload core.ContainerPayload

	err := json.Unmarshal(msgBytes, &payload)
	if err != nil {
		log.ErrorF("Error decode json message: %s", err.Error())
		return err
	}
	log.Debugf("stared provider container payload:[%v] %+v", payload.ID, payload.Name)

	resultMsg, err := c.provider.Execute(ctx, "create", payload)
	if err != nil {
		log.ErrorF("Error create container: %s", err.Error())
		return err
	}

	event := core.Event{
		ID:         core.GenerateID(),
		ResourceID: payload.ID,
		Type:       core.ContainerStarted,
		Payload:    msgBytes,
		CreatedAt:  time.Now(),
	}
	err = c.repository.Save(ctx, event)
	if err != nil {
		log.ErrorF("Error saving ContainerStarted event to database: %s", err.Error())
		return err
	}

	log.InforF("Created container: %s", resultMsg)
	return nil
}
