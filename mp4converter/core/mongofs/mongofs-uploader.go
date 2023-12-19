package mongofs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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
