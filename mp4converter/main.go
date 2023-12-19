package main

import (
	"log"

	"github.com/FadyGamilM/mp4converter/core/rabbitmq"
)

const (
	exchangeName    = "audio_upload"
	audioRoutingKey = "audio_uploaded"
	videoRoutingKey = "video_uploaded"
	queueName       = "video_converter"
	contentType     = "application/json"
	converterQueue  = "audio_converter"
)

func main() {

	client, err := rabbitmq.Connect("fady", "fady", "localhost:5672", "video_converter")

	if err != nil {
		log.Fatalf("error : %v", err)
	}

	defer client.Conn.Close()
	defer client.Ch.Close()

	messageBus, err := client.Consume("video", "9", false)
	if err != nil {
		log.Fatalf("error : %v", err)
	}

	var blocking chan struct{}

	go func() {
		for message := range messageBus {
			// breakpoint here
			log.Printf("New Message: %v", message)
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")
	// This will block forever
	<-blocking
}
