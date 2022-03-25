package users

import (
	"github.com/oh-jinsu/helloworld/entities"
	"github.com/oh-jinsu/helloworld/models"
	"gorm.io/gorm"
)

func usernameExists(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&models.User{}).Error

	return err == nil
}

func saveUser(db *gorm.DB, entity *entities.User) *entities.User {
	db.Create(&models.User{
		Username: entity.Username().ToString(),
		Password: entity.Password().ToString(),
	})

	result := &models.User{}

	db.Where("username = ?", entity.Username().ToString()).First(result)

	return entities.NewUser(
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
	)
}
