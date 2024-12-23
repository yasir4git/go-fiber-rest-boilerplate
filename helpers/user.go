package helpers

import (
	"errors"

	"github.com/IKHINtech/go-fiber-rest-boilerplate/app/models"
	"github.com/IKHINtech/go-fiber-rest-boilerplate/database"
	"gorm.io/gorm"
)

func GetUserByEmail(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Email: &e}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByPhone(p string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where(&models.User{Phone: p}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
