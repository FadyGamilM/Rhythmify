package converter

import (
	"io"
	"os"
	"os/exec"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

// ConvertMP4toMP3 converts an MP4 file to an MP3 file
func ConvertMP4toMP3(inputFile, outputFile string) error {
	// Use FFmpeg to convert MP4 to WAV
	wavFile := "temp.wav"
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-acodec", "pcm_s16le", "-ar", "44100", wavFile)
	if err := cmd.Run(); err != nil {
		return err
	}

	// Convert WAV to MP3
	if err := convertWAVtoMP3(wavFile, outputFile); err != nil {
		return err
	}

	// Remove temporary WAV file
	if err := os.Remove(wavFile); err != nil {
		return err
	}

	return nil
}

// convertWAVtoMP3 converts a WAV file to an MP3 file
func convertWAVtoMP3(inputFile, outputFile string) error {
	wavFile, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer wavFile.Close()

	mp3File, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	// Convert WAV to MP3
	if err := mp3.Encode(mp3File, wavFile, &mp3.Options{BitRate: 128}); err != nil {
		return err
	}

	return nil
}

// PlayMP3 plays an MP3 file
func PlayMP3(file string) error {
	mp3File, err := os.Open(file)
	if err != nil {
		return err
	}
	defer mp3File.Close()

	p, err := oto.NewPlayer(44100, 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	_, err = io.Copy(p, mp3File)
	if err != nil {
		return err
	}

	return nil
}
