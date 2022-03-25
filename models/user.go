package models

import (
	"github.com/oh-jinsu/helloworld/entities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func UsernameExists(db *gorm.DB, username string) bool {
	err := db.Where("username = ?", username).First(&User{}).Error

	return err == nil
}

func SaveUser(db *gorm.DB, username string, password string) {
	db.Create(&User{
		Username: username,
		Password: password,
	})
}

func FindUserByUsername(db *gorm.DB, username string) (*entities.User, error) {
	result := &User{}

	if err := db.Where("username = ?", username).First(result).Error; err != nil {
		return &entities.User{}, err
	}

	return &entities.User{
		Id:        result.ID,
		Username:  &entities.Username{Value: result.Username},
		Password:  &entities.Password{Value: result.Password},
		CreatedAt: result.CreatedAt,
	}, nil
}
