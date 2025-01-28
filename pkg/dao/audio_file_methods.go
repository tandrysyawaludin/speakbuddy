package dao

import (
	"log"
	"speakbuddy-be/pkg/dto"
)

var audioFileOrm = NewAudiofileOrm()

func GetAudioFile(audioFileParam *dto.AudioFile) (*dto.AudioFile, error) {
	// get audio file
	audioFile, err := audioFileOrm.Get(audioFileParam)
	if err != nil {
		log.Printf("[ERROR] get audio file failed, err: %+v", err)
		return nil, err
	}

	return audioFile, nil
}

func PostAudioFile(newAudioFile *dto.AudioFile) error {
	// post audio file
	err := audioFileOrm.Post(newAudioFile)
	if err != nil {
		log.Printf("[ERROR] post audio file failed, err: %+v", err)
		return err
	}

	return nil
}
