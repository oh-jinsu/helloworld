package models

import (
	"time"

	"github.com/oh-jinsu/helloworld/entities"
	"gorm.io/gorm"
)

type User struct {
	Id           uint `gorm:"primaryKey"`
	Username     string
	Password     string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func NextUserId(db *gorm.DB) uint {
	user := &User{}

	db.Last(user)

	return user.Id + 1
}

func ExistsByUsername(db *gorm.DB, username *entities.Username) bool {
	err := db.Where("username = ?", username.ToString()).First(&User{}).Error

	return err == nil
}

func SaveUser(db *gorm.DB, user *entities.User) {
	db.Save(&User{
		Id:           user.Id(),
		Username:     user.Username().ToString(),
		Password:     user.Password().ToString(),
		RefreshToken: user.RefreshToken(),
		CreatedAt:    user.CreatedAt(),
	})
}

func FindUserById(db *gorm.DB, id uint) (*entities.User, error) {
	result := &User{}

	if err := db.Where("id = ?", id).First(result).Error; err != nil {
		return &entities.User{}, err
	}

	return entities.NewUser(
		result.Id,
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
		result.RefreshToken,
		result.CreatedAt,
	), nil
}

func FindUserByUsername(db *gorm.DB, username *entities.Username) (*entities.User, error) {
	result := &User{}

	if err := db.Where("username = ?", username.ToString()).First(result).Error; err != nil {
		return &entities.User{}, err
	}

	return entities.NewUser(
		result.Id,
		entities.NewUsername(result.Username),
		entities.NewPassword(result.Password),
		result.RefreshToken,
		result.CreatedAt,
	), nil
}
