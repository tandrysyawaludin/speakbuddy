package dao

import (
	"speakbuddy-be/pkg/dto"

	"gorm.io/gorm"
)

type userDaoInterface interface {
	Get(userParam *dto.User) (*dto.User, error)
	GetById(userId int) (*dto.User, error)
	Post(user *dto.User) error
}

type userApiOrm struct {
}

func NewUserOrm() userDaoInterface {
	return &userApiOrm{}
}

func (m *userApiOrm) Get(userParam *dto.User) (*dto.User, error) {
	var user *dto.User
	if err := daoConfig.Model(&dto.User{}).Where(userParam).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *userApiOrm) GetById(userId int) (*dto.User, error) {
	var user *dto.User
	if err := daoConfig.Model(&dto.User{}).First(&user, userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *userApiOrm) Post(user *dto.User) error {
	if err := daoConfig.Session(&gorm.Session{FullSaveAssociations: true, CreateBatchSize: 100}).Model(dto.User{}).Create(user).Error; err != nil {
		return err
	}

	return nil
}
