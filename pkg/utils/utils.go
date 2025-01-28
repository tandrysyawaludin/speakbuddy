package utils

import (
	"errors"
	"log"
	"os/exec"
)

// ConvertMp3ToWav converts an MP3 file to WAV format using FFmpeg
func ConvertMp3ToWav(inputPath, outputPath string) error {
	// Check if input and output paths are provided
	if inputPath == "" || outputPath == "" {
		return errors.New("input and output paths must not be empty")
	}

	// FFmpeg command to convert MP3 to WAV
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ar", "44100", "-ac", "2", "-f", "wav", outputPath)

	// Run the command and capture any errors
	if err := cmd.Run(); err != nil {
		log.Printf("[ERROR] convert mp3 to wav failed, %s %s err: %+v", inputPath, outputPath, err,)
		return errors.New("failed to convert mp3 to wav")
	}

	return nil
}

func ConvertWavToMp3(inputPath, outputPath string) error {
	// Check if input and output paths are provided
	if inputPath == "" || outputPath == "" {
		return errors.New("input and output paths must not be empty")
	}

	// FFmpeg command to convert WAV to MP3
	cmd := exec.Command("ffmpeg", "-i", inputPath, "-codec:a", "libmp3lame", "-qscale:a", "2", outputPath)

	// Run the command and capture any errors
	if err := cmd.Run(); err != nil {
		log.Printf("[ERROR] convert wav to mp3 failed, err: %+v", err)
		return errors.New("failed to convert wav to mp3")
	}

	return nil
}
