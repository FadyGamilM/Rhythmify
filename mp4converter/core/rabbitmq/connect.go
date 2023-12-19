package rabbitmq

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitClient is used to keep track of the RabbitMQ connection
type RabbitClient struct {
	// The connection that is used
	Conn *amqp.Connection
	// The channel that processes/sends Messages
	Ch *amqp.Channel
}

// ConnectRabbitMQ will spawn a Connection
func Connect(username, password, host, vhost string) (*RabbitClient, error) {
	// Setup the Connection to RabbitMQ host using AMQP
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
	if err != nil {
		return nil, err
	}
	// Unique, Conncurrent Server Channel to process/send messages
	// A good rule of thumb is to always REUSE Conn across applications
	// But spawn a new Channel per routine
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitClient{
		Conn: conn,
		Ch:   ch,
	}, nil
}
