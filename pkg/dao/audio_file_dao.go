package dao

import (
	"speakbuddy-be/pkg/dto"

	"gorm.io/gorm"
)

type audiofileDaoInterface interface {
	Get(audioFileParam *dto.AudioFile) (*dto.AudioFile, error)
	Post(audioFile *dto.AudioFile) error
}

type audiofileApiOrm struct {
}

func NewAudiofileOrm() audiofileDaoInterface {
	return &audiofileApiOrm{}
}

func (m *audiofileApiOrm) Get(audioFileParam *dto.AudioFile) (*dto.AudioFile, error) {
	var audioFile *dto.AudioFile
	if err := daoConfig.Model(&dto.AudioFile{}).Where(audioFileParam).Find(&audioFile).Error; err != nil {
		return nil, err
	}

	return audioFile, nil
}

func (m *audiofileApiOrm) Post(audioFile *dto.AudioFile) error {
	if err := daoConfig.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: 100}).Model(dto.AudioFile{}).Create(audioFile).Error; err != nil {
		return err
	}

	return nil
}
