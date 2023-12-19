package rabbitmq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConsumer(t *testing.T) {
	test := struct {
		queueName string
		consumer  string
	}{
		queueName: "video",
		consumer:  "5",
	}

	client, err := Connect("fady", "fady", "localhost:5672", "video_converter")

	assert.NoError(t, err)
	defer client.Conn.Close()
	defer client.Ch.Close()

	messageBus, err := client.Consume(test.queueName, test.consumer, false)
	assert.NoError(t, err)

	for message := range messageBus {
		// breakpoint here
		t.Logf("new Message: %v", message)
	}

}
