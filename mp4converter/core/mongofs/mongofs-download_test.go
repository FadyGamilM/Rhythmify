package mongofs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const mp4ID = "6581b71c023f772a2d820794"

func TestDownloadMp4File(t *testing.T) {
	testCase := struct {
		fileID string
	}{
		fileID: mp4ID,
	}
	mongoFS, err := Connect()
	// if err != nil {
	// 	return fmt.Errorf("error connecting to mongodb instance : %v\n", err)
	// }
	assert.NoError(t, err)

	fileid, err := primitive.ObjectIDFromHex(testCase.fileID)
	assert.NoError(t, err)
	err = DownloadFileByID(mongoFS, fileid, "output.mp4")
	assert.NoError(t, err)
}
