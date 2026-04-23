package rabbitmq

import (
	"context"
	"encoding/json"

	"github.com/Aaron-GMM/DockOps/internal/config/logger"
	amqp "github.com/rabbitmq/amqp091-go"
)

var log = logger.NewLogger("RabbitMQ-Publisher")

type RabbitPublisher struct {
	Channel *amqp.Channel
}

func NewRabbitPublisher(conn Connection) *RabbitPublisher {
	return &RabbitPublisher{
		Channel: conn.Ch,
	}
}

func (p *RabbitPublisher) Publish(ctx context.Context,
	queueName string, message []byte) error {
	log.Debugf("Preparete publishing de %d bytes for queue: %s", len(message), queueName)
	err := p.Channel.PublishWithContext(
		ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         message,
		},
	)
	if err != nil {
		log.ErrorF("Broker failed to publish: %s", err)
		return err
	}
	log.InforF("Publishing de %d bytes for queue: %s", len(message), queueName)
	return nil
}

func (p *RabbitPublisher) PublishEvent(ctx context.Context,
	exchange string, routingKey string, event any) error {
	log.Debugf("Preparete publishing de %d bytes for event: %s", len(event.([]byte)), exchange)
	body, err := json.Marshal(event)

	if err != nil {
		log.ErrorF("Broker failed to marshal event: %s", err)
		return err
	}
	log.Debugf("Publishing de %d bytes for event: %s", len(body), exchange)
	return p.Channel.PublishWithContext(
		ctx,
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
