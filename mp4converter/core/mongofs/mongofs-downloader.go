package mongofs

import (
	"fmt"
	"io"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

// Download a file from GridFS using the document ID
func DownloadFileByID(mongoFS *gridfs.Bucket, fileID primitive.ObjectID, destination string) error {

	downloadStream, err := mongoFS.OpenDownloadStream(fileID)
	if err != nil {
		return fmt.Errorf("error trying to open a stream to download the file : %v\n", err)
	}
	defer downloadStream.Close()

	outputFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error trying to open output destination to download the file into it : %v\n", err)
	}
	defer outputFile.Close()

	// Copy the file content to the output file
	_, err = io.Copy(outputFile, downloadStream)
	if err != nil {
		return fmt.Errorf("error copying file content : %v", err)
	}

	fmt.Printf("File downloaded successfully to %v\n", outputFile)
	return nil
}
