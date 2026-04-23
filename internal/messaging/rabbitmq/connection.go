package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

type Connection struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewConnection(url string) (*Connection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &Connection{
		Conn: conn,
		Ch:   ch,
	}, nil
}

func (conn *Connection) Close() {
	if conn.Conn != nil {
		conn.Conn.Close()
	}
	if conn.Ch != nil {
		conn.Ch.Close()
	}
}
