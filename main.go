package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/oh-jinsu/helloworld/models"
	"github.com/oh-jinsu/helloworld/modules/app"
	"github.com/oh-jinsu/helloworld/modules/users"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("I failed to load the .env file")
	}

	dsn := os.Getenv("DSN")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("I failed to connect the database")
	}

	db.AutoMigrate(&models.User{})

	router := gin.Default()

	appModule := &app.Module{
		Router: router.Group(""),
	}

	appModule.AddWelcomeController()

	userModule := &users.Module{
		Router: router.Group("users"),
		DB:     db,
	}

	userModule.AddCreateUserController()

	mode := os.Getenv("MODE")

	gin.SetMode(mode)

	addr := fmt.Sprintf("localhost:%s", os.Getenv("PORT"))

	router.Run(addr)
}
