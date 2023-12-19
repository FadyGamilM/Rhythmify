package converter

import (
	"fmt"
	"os/exec"
)

func ConvertMP4toMP3(inputPath string, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vn", "-acodec", "libmp3lame", "-ac", "2", "-q:a", "4", "-y", outputPath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error converting MP4 to MP3: %v", err)
	}

	fmt.Printf("File converted successfully to %s\n", outputPath)
	return nil
}
