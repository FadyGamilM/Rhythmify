package mongogridfs

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	db := client.Database("mydb")

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

// Upload a file to GridFS and return the document ID
func UploadFile(fs *gridfs.Bucket, filename string, content []byte) (primitive.ObjectID, error) {

	uploadStream, err := fs.OpenUploadStream(
		filename,
		options.GridFSUpload().SetMetadata(map[string]string{"type": "text"}),
	)
	if err != nil {
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(content)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return uploadStream.FileID.(primitive.ObjectID), nil
}

// Download a file from GridFS using the document ID
func DownloadFile(fs *gridfs.Bucket, fileID primitive.ObjectID, destination string) error {
	downloadStream, err := fs.OpenDownloadStream(fileID)
	if err != nil {
		return err
	}
	defer downloadStream.Close()

	data, err := ioutil.ReadAll(downloadStream)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destination, data, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("File with ID '%s' downloaded to '%s'\n", fileID.Hex(), destination)

	return nil
}
