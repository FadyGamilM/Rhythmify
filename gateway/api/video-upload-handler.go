package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	mongogridfs "github.com/FadyGamilM/rhythmify/gateway/mongo-gridfs"
	"github.com/FadyGamilM/rhythmify/gateway/rabbitmq"
	videoevents "github.com/FadyGamilM/rhythmify/gateway/video-events"
	"github.com/gin-gonic/gin"
)

const (
	exchangeName = "video_upload"
	queueName    = "video"
	routingKey   = "video_uploaded"
	contentType  = "application/json"
)

func (h *Handler) UploadVideo(c *gin.Context) {
	// Get the uploaded file
	_, file, err := c.Request.FormFile("file") // or "files" depending on your form field name
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Open the file
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer openedFile.Close()

	// ======== set to mongo fsgrid and after that configure and publish to rabbitmq
	// Read the content of the video file into a byte slice
	videoContent, err := ioutil.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error reading video content"})
		return
	}

	fileID, err := mongogridfs.UploadFile(h.mongoFS, file.Filename, videoContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error uploading video content"})
		return
	}

	// ======== rabbitmq connection
	conn, err := rabbitmq.ConnectRabbitMQ("fady", "fady", "localhost:5672", "video_converter")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error connecting to rabbitmq .. ",
		})
	}

	defer conn.Close()
	client, err := rabbitmq.NewRabbitMQClient(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error creating new channel to rabbitmq connection.. ",
		})
	}
	defer client.Close()

	userID, _ := c.MustGet("userId").(int64)
	userEmail, _ := c.MustGet("email").(string)
	EventMsg := videoevents.UploadVideoEvent{
		VideoFileId: fileID.Hex(),
		AudioFileId: "",
		UserId:      userID,
		Email:       userEmail,
	}

	// Convert the EventMsg struct to a JSON string
	eventMsgJSON, err := json.Marshal(EventMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error marshaling event message to JSON",
		})
		return
	}

	err = client.ProduceMessage(exchangeName, routingKey, contentType, []byte(eventMsgJSON))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error producing message event to the queue",
		})
		return
	}
	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully " + file.Filename})
}
