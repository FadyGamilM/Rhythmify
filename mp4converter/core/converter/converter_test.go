package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverter(t *testing.T) {
	testCase := struct {
		inputPath  string
		outputPath string
	}{
		inputPath:  "output.mp4",
		outputPath: "output.mp3",
	}

	err := ConvertMP4toMP3(testCase.inputPath, testCase.outputPath)
	assert.NoError(t, err)
}
