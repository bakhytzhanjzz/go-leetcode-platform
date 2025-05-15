package database

import (
	"fmt"
	"log"

	"github.com/bakhytzhanjzz/go-leetcode-platform/user-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=postgres dbname=user_service port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("DB migration error:", err)
	}
	fmt.Println("Connected to user DB")
}
