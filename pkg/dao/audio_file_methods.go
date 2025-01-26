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
		log.Print("error Getting audio file from DB", err.Error())
		return nil, err
	}

	return audioFile, nil
}

func PostAudioFile(newAudioFile *dto.AudioFile) error {
	// post audio file
	err := audioFileOrm.Post(newAudioFile)
	if err != nil {
		log.Print("error posting audio file to DB", err.Error())
		return err
	}

	return nil
}
