package dto

import "gorm.io/gorm"

type AudioFile struct {
	gorm.Model
	UserId   int    `json:"user_id"`
	PhraseId int    `json:"phrase_id"`
	FilePath string `json:"file_path"`
}
