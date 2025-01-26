package speakbuddybeapp

import (
	"speakbuddy-be/pkg/dto"

	"gorm.io/gorm"
)

func DbInit(db *gorm.DB) error {
	// Migrate the schema
	if err := db.AutoMigrate(&dto.AudioFile{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&dto.Phrase{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&dto.User{}); err != nil {
		return err
	}
	return nil
}
