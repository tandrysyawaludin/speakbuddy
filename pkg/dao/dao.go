package dao

import "gorm.io/gorm"

var daoConfig *gorm.DB

func Init(newDAOConfig *gorm.DB) {
	daoConfig = newDAOConfig
}
