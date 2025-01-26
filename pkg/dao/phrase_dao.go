package dao

import (
	"speakbuddy-be/pkg/dto"

	"gorm.io/gorm"
)

type phraseDaoInterface interface {
	Get(phraseParam *dto.Phrase) (*dto.Phrase, error)
	GetById(phraseId int) (*dto.Phrase, error)
	Post(phrase *dto.Phrase) error
}

type phraseApiOrm struct {
}

func NewPhraseOrm() phraseDaoInterface {
	return &phraseApiOrm{}
}

func (m *phraseApiOrm) Get(phraseParam *dto.Phrase) (*dto.Phrase, error) {
	var phrase *dto.Phrase
	if err := daoConfig.Model(&dto.Phrase{}).Where(phraseParam).Find(&phrase).Error; err != nil {
		return nil, err
	}

	return phrase, nil
}

func (m *phraseApiOrm) GetById(phraseId int) (*dto.Phrase, error) {
	var phrase *dto.Phrase
	if err := daoConfig.Model(&dto.Phrase{}).First(&phrase, phraseId).Error; err != nil {
		return nil, err
	}

	return phrase, nil
}

func (m *phraseApiOrm) Post(phrase *dto.Phrase) error {
	if err := daoConfig.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: 100}).Model(dto.Phrase{}).Create(phrase).Error; err != nil {
		return err
	}

	return nil
}
