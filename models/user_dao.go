package models

import (
	"github.com/oh-jinsu/helloworld/entities"
	"gorm.io/gorm"
)

func UsernameExists(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&User{}).Error

	return err == nil
}

func SaveUser(db *gorm.DB, entity *entities.User) {
	db.Create(&User{
		Username: entity.Username().ToString(),
		Password: entity.Password().ToString(),
	})
}

func FindUserByUsername(db *gorm.DB, username *entities.Username) (*entities.User, error) {
	result := &User{}

	if err := db.Where("username = ?", username.ToString()).First(result).Error; err != nil {
		return &entities.User{}, err
	}

	return entities.NewUser(
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
	), nil
}
