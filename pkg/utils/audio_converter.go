package utils

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func ConvertMp3ToWav(inputPath, outputPath string) error {
	if inputPath == "" || outputPath == "" {
		return errors.New("input and output paths must not be empty")
	}

	os.Remove(outputPath)
	defer os.Remove(inputPath)

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-ar", "44100", "-ac", "2", "-f", "wav", outputPath)

	if err := cmd.Run(); err != nil {
		log.Printf("[ERROR] convert mp3 to wav failed, %s %s err: %+v", inputPath, outputPath, err)
		return errors.New("failed to convert mp3 to wav")
	}

	if err := UploadToSftp(outputPath, outputPath); err != nil {
		return err
	}

	return nil
}

func ConvertWavToMp3(outputPath, inputPath string) error {
	if inputPath == "" || outputPath == "" {
		return errors.New("input and output paths must not be empty")
	}

	os.Remove(outputPath)
	os.Remove(inputPath)

	if err := DownloadFromSftp(outputPath, inputPath); err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-codec:a", "libmp3lame", "-qscale:a", "2", outputPath)

	if err := cmd.Run(); err != nil {
		log.Printf("[ERROR] convert wav to mp3 failed, err: %+v", err)
		return errors.New("failed to convert wav to mp3")
	}

	return nil
}
