package dao

import (
	"log"
	"speakbuddy-be/pkg/dto"
)

var userOrm = NewUserOrm()

func GetUserById(userId int) (*dto.User, error) {
	user, err := userOrm.GetById(userId)
	if err != nil {
		log.Print("error Getting user file from DB", err.Error())
		return nil, err
	}

	return user, nil
}
