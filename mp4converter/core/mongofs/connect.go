package mongofs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*gridfs.Bucket, error) {
	// MongoDB connection string
	connectionString := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Printf("error trying to connect to mongo : %v\n", err)
		return nil, err
	}

	// Access the "mydb" database
	db := client.Database("mp4")

	// Create a GridFS bucket
	fs, err := gridfs.NewBucket(
		db,
		options.GridFSBucket().SetName("fs"),
	)
	if err != nil {
		log.Printf("error trying to create a mongodb grid fs : %v\n", err)
		return nil, err
	}

	return fs, nil
}
