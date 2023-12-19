package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

// ProduceMessage sends a message to the specified exchange with the given routing key
func (rc *RabbitClient) ProduceMessage(exchange, routingKey, contentType string, body []byte) error {

	// Publish the message to the exchange
	err := rc.Ch.Publish(
		exchange,   // Exchange name
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Println("Message published successfully!")
	return nil
}
