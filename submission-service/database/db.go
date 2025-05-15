package database

import (
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/submission-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=submissions port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	err = db.AutoMigrate(&models.Submission{})
	if err != nil {
		log.Fatalf("DB migration error: %v", err)
	}

	DB = db
	fmt.Println("Connected to DB")
}
