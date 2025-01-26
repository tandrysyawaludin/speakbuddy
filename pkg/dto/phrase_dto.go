package dto

import "gorm.io/gorm"

type Phrase struct {
	gorm.Model
	Phrase string `json:"phrase"`
}
