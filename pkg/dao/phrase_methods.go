package dao

import (
	"log"
	"speakbuddy-be/pkg/dto"
)

var phraseOrm = NewPhraseOrm()

func GetPhraseById(phraseId int) (*dto.Phrase, error) {
	// get phrase file
	phrase, err := phraseOrm.GetById(phraseId)
	if err != nil {
		log.Print("error getting phrase file from db", err.Error())
		return nil, err
	}

	return phrase, nil
}
