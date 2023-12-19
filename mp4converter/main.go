package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	mongogridfs "github.com/FadyGamilM/mp4converter/mongo-gridfs"
	"github.com/FadyGamilM/mp4converter/rabbitmq"
	videoevents "github.com/FadyGamilM/mp4converter/video-events"
	"github.com/streadway/amqp"
)

const (
	exchangeName   = "video_upload"
	routingKey     = "video_uploaded"
	queueName      = "video_converter"
	contentType    = "application/json"
	converterQueue = "audio_converter"
)

func main() {
	// Connect to MongoDB GridFS
	fs, err := mongogridfs.Connect()
	if err != nil {
		log.Fatal("Error connecting to MongoDB GridFS:", err)
	}

	// Connect to RabbitMQ
	conn, err := rabbitmq.ConnectRabbitMQ("fady", "fady", "localhost:5672", "video_converter")
	if err != nil {
		log.Fatal("Error connecting to RabbitMQ:", err)
	}
	defer conn.Close()

	// Create RabbitMQ client
	client, err := rabbitmq.NewRabbitMQClient(conn)
	if err != nil {
		log.Fatal("Error creating RabbitMQ client:", err)
	}
	defer client.Close()

	// Set up a signal handler to gracefully shutdown the service
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChannel
		log.Printf("Received signal %v. Shutting down...\n", sig)
		// Perform cleanup or additional shutdown tasks here
		os.Exit(0)
	}()

	// Start consuming messages from RabbitMQ
	messages, err = client.Consume(queueName, "", handleVideoUploadedEvent)
	if err != nil {
		log.Fatal("Error consuming RabbitMQ messages:", err)
	}

	messageBus, err := client.Consume("customers_created", "email-service", false)
	if err != nil {
		panic(err)
	}

	// blocking is used to block forever
	var blocking chan struct{}

	go func() {
		for message := range messageBus {
			// breakpoint here
			log.Printf("New Message: %v", message)
			// Multiple means that we acknowledge a batch of messages, leave false for now
			if err := message.Ack(false); err != nil {
				log.Printf("Acknowledged message failed: Retry ? Handle manually %s\n", message.MessageId)
				continue
			}
			log.Printf("Acknowledged message %s\n", message.MessageId)
		}
	}()

	log.Println("Consuming, to close the program press CTRL+C")
	// This will block forever
	<-blocking

}

func handleVideoUploadedEvent(delivery amqp.Delivery) {
	// Decode the JSON message
	var event videoevents.UploadVideoEvent
	err := json.Unmarshal(delivery.Body, &event)
	if err != nil {
		log.Println("Error decoding JSON message:", err)
		return
	}

	log.Printf("Received video uploaded event. VideoFileId: %s\n", event.VideoFileId)

	// Download the MP4 file from MongoDB GridFS
	mp4Content, err := mongogridfs.DownloadFileByID(event.VideoFileId)
	if err != nil {
		log.Println("Error downloading MP4 file from MongoDB GridFS:", err)
		return
	}

	// Convert the MP4 content to MP3 (you can use the provided ConvertMP4toMP3 function)
	mp3Content, err := ConvertMP4toMP3(mp4Content)
	if err != nil {
		log.Println("Error converting MP4 to MP3:", err)
		return
	}

	// Upload the MP3 content to MongoDB GridFS
	mp3FileID, err := mongogridfs.UploadFile(fs, "output.mp3", mp3Content)
	if err != nil {
		log.Println("Error uploading MP3 file to MongoDB GridFS:", err)
		return
	}

	log.Printf("MP3 file converted and uploaded. MP3 File ID: %s\n", mp3FileID.Hex())

	// Publish an event indicating the completion of the conversion
	mp3Event := videoevents.ConvertVideoEvent{
		VideoFileId: event.VideoFileId,
		AudioFileId: mp3FileID.Hex(),
	}
	mp3EventJSON, err := json.Marshal(mp3Event)
	if err != nil {
		log.Println("Error marshaling MP3 event to JSON:", err)
		return
	}

	err = client.ProduceMessage(exchangeName, converterQueue, contentType, mp3EventJSON)
	if err != nil {
		log.Println("Error publishing MP3 event to RabbitMQ:", err)
		return
	}

	log.Println("MP3 conversion event published successfully!")
}
