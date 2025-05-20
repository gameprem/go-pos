package database

import (
	"log"

	"go-pos/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.UserInfo{})
}
