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

func UsernameExists(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&User{}).Error

	return err == nil
}

func SaveUser(db *gorm.DB, entity *entities.User) {
	db.Create(&User{
		Username: entity.Username.ToString(),
		Password: entity.Password.ToString(),
	})
}

func FindUser(db *gorm.DB, user *entities.User) error {
	result := &User{}

	if err := db.Where("username = ?", user.Username.ToString()).First(result).Error; err != nil {
		return err
	}

	user.Id = result.ID

	user.Username = &entities.Username{Value: result.Username}
	user.Password = &entities.Password{Value: result.Password}
	user.CreatedAt = result.CreatedAt

	return nil
}
