package core

import "context"

type MessagePublisher interface {
	Publish(ctx context.Context, queueName string, message []byte) error
}
