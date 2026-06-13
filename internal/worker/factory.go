package worker

import (
	"context"

	"github.com/Aaron-GMM/DockOps/internal/messaging/rabbitmq"
	"github.com/Aaron-GMM/DockOps/internal/provider/docker"
	"github.com/Aaron-GMM/DockOps/internal/storage/postgres"
	"gorm.io/gorm"
)

func StartWorker(ctx context.Context, db *gorm.DB, rabbitConn rabbitmq.Connection) error {
	repo := postgres.NewEventRepository(db)
	dockerProvider, err := docker.NewDockerClient()
	if err != nil {
		return err
	}

	consumer := NewContainerConsumer(dockerProvider, repo)
	subscriber := rabbitmq.NewRabbitSubscriber(rabbitConn)

	return subscriber.Subscribe(ctx, "container_tasks", consumer.ProcessMessage)
}
