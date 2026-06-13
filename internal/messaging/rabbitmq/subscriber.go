package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitSubscriber struct {
	Channel *amqp.Channel
}

func NewRabbitSubscriber(conn Connection) *RabbitSubscriber {
	return &RabbitSubscriber{
		Channel: conn.Ch,
	}
}

func (s *RabbitSubscriber) Subscribe(ctx context.Context, queueName string, handler func(ctx context.Context, body []byte) error) error {
	_, err := s.Channel.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := s.Channel.Consume(
		queueName,
		"",    // consumer
		false, // auto-ack (we want manual ack)
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			err := handler(ctx, d.Body)
			if err != nil {
				log.ErrorF("Error processing message: %s", err)
				// Nack and requeue
				d.Nack(false, true)
			} else {
				d.Ack(false)
			}
		}
	}()

	return nil
}
