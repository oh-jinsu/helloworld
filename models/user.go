package models

import (
	"github.com/oh-jinsu/helloworld/entities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	Password     string
	RefreshToken string
}

func UsernameExists(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&User{}).Error

	return err == nil
}

func PutUser(db *gorm.DB, user *entities.User) {
	if user.Id() != 0 {
		db.Save(&User{
			Username:     user.Username().ToString(),
			Password:     user.Password().ToString(),
			RefreshToken: user.RefreshTOken(),
		})
	}

	db.Create(&User{
		Username:     user.Username().ToString(),
		Password:     user.Password().ToString(),
		RefreshToken: user.RefreshTOken(),
	})
}

func FindUserByUsername(db *gorm.DB, username *entities.Username) (*entities.User, error) {
	result := &User{}

	if err := db.Where("username = ?", username.ToString()).First(result).Error; err != nil {
		return &entities.User{}, err
	}

	return entities.NewUser(
		result.ID,
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
		"",
		result.CreatedAt,
	), nil
}
