package repositories

import (
	"github.com/oh-jinsu/helloworld/entities"
	"github.com/oh-jinsu/helloworld/models"
	"gorm.io/gorm"
)

func UsernameExists(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&models.User{}).Error

	return err == nil
}

func SaveUser(db *gorm.DB, entity *entities.User) {
	db.Create(&models.User{
		Username: entity.Username().ToString(),
		Password: entity.Password().ToString(),
	})
}

func FindUserByUsername(db *gorm.DB, username *entities.Username) *entities.User {
	result := &models.User{}

	db.Where("username = ?", username.ToString()).First(result)

	return entities.NewUser(
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
	)
}
