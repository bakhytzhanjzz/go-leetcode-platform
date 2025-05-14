package database

import (
	"fmt"
	"github.com/bakhytzhanjzz/go-leetcode-platform/problem-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=problems port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.Category{}, &models.Problem{})
	if err != nil {
		log.Fatal("DB migration error:", err)
	}

	fmt.Println("Connected to DB")
	return db
}
