package rabbitmq

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitPublisher struct {
	Channel *amqp.Channel
}

func NewRabbitPublisher(conn Connection) *RabbitPublisher {
	return &RabbitPublisher{
		Channel: conn.Ch,
	}
}

func (p *RabbitPublisher) Publish(exchange string,
	routingKey string, msgBytes []byte) error {

	return p.Channel.PublishWithContext(
		context.Background(),
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgBytes,
		},
	)
}
func (p *RabbitPublisher) PublishEvent(exchange string,
	rountingKey string, event any) error {
	body, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.Publish(exchange, rountingKey, body)
}
