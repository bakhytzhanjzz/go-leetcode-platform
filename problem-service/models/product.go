package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"` // Easy, Medium, Hard
	CategoryID  uint   `json:"category_id"`
}
